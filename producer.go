package factogo

import (
	"reflect"
)

func (this *FactoryInstance) Produce(object interface{}) error {
	modifiedFactory := this
	originalFactory := factories[this.name]

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

	return nil
}
