package factogo

import (
	"reflect"
)

func (this *FactoryInstance) Produce(object interface{}) error {
	modifiedFactory := this
	originalFactory := factories[this.name]
	if !modifiedFactory.notPersist && modifiedFactory.persistFunction == nil {
		modifiedFactory.persistFunction = originalFactory.persistFunction
		modifiedFactory.notPersist = originalFactory.notPersist
	}

	if !modifiedFactory.notPersist && modifiedFactory.afterPersistFunction == nil {
		modifiedFactory.afterPersistFunction = originalFactory.afterPersistFunction
	}

	for _, modifiedValue := range modifiedFactory.values {
		originalFactory.values[modifiedValue.name].value = modifiedValue.value
	}

	for _, value := range originalFactory.values {
		valueType := reflect.TypeOf(value.value)
		field := reflect.Indirect(reflect.ValueOf(object)).FieldByName(value.name)

		if valueType.Kind() == reflect.Func {

			if reflect.TypeOf(value.value).NumOut() > 0 {
				_value := reflect.ValueOf(value.value).Call([]reflect.Value{})[0].Interface()
				field.Set(reflect.ValueOf(_value))
			} else {
				reflect.ValueOf(value.value).Call([]reflect.Value{reflect.ValueOf(object)})
			}
		} else {
			field.Set(reflect.ValueOf(value.value))
		}
	}

	if !this.notPersist && (this.persistFunction != nil || persistFunction != nil) {
		if this.persistFunction == nil {
			this.persistFunction = persistFunction
		}
		this.persistFunction(object)

		if this.afterPersistFunction == nil {
			this.afterPersistFunction = afterPersistFunction
		}
		if this.afterPersistFunction != nil {
			this.afterPersistFunction(object)
		}
	}

	return nil
}
