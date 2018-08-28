package factogo

import (
	"reflect"
)

/*
Produce invokes a designed Factory to create a new product Struct.

The values for the product Struct are copied in the target parameter passed
to the Produce function.

Example:
	Factory("staff").Design(Staff{})
	target := &Staff{}
	Factory("staff").Produce(target)
*/
func (this *factoryInstance) Produce(object interface{}) error {
	modifiedFactory := this
	originalFactory := factories[this.name]
	if !modifiedFactory.notPersist && modifiedFactory.persistFunction == nil {
		modifiedFactory.persistFunction = originalFactory.persistFunction
		modifiedFactory.notPersist = originalFactory.notPersist
	}

	if !modifiedFactory.notPersist && modifiedFactory.afterPersistFunction == nil {
		modifiedFactory.afterPersistFunction = originalFactory.afterPersistFunction
	}

	for _, originalValue := range originalFactory.values {
		if _, ok := modifiedFactory.values[originalValue.name]; !ok {
			modifiedFactory.values[originalValue.name] = &factoryValue{
				name:   originalValue.name,
				value:  originalValue.value,
				params: originalValue.params,
			}
		}
	}

	for _, value := range modifiedFactory.values {
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
