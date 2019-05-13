package factogo

import (
	"errors"
	"fmt"
	"reflect"
)

/*
Design ends the process to design a Factory instance. When a Factory instance is
designed then it is registered and can be produced calling the method Produce().

Example:

	Factory("staff").Design(&Staff{}) // Design the Factory Instance
	Factory("staff").Produce() // Now is possible to produce from the designed Factory
*/
func (this *factoryInstance) Design(object interface{}) error {
	err := checkValues(object, this)
	if err != nil {
		return err
	}
	factories[this.name] = this
	return nil
}

func checkValues(object interface{}, factory *factoryInstance) error {

	objectValue := reflect.ValueOf(object)
	objectType := reflect.Indirect(objectValue).Type()
	for _, _factoryValue := range factory.values {

		fieldName := _factoryValue.name

		field, ok := objectType.FieldByName(fieldName)

		if !ok {
			return errors.New(fmt.Sprintf(
				"The struct '%v' doesn't have a field '%v'",
				objectType.Name(),
				fieldName))
		}

		if _factoryValue.value != nil {
			value := reflect.TypeOf(_factoryValue.value)

			var targetKind reflect.Kind

			if value.Kind() == reflect.Func {
				targetKind = value.Out(0).Kind()
			} else {
				targetKind = value.Kind()
			}

			if field.Type.Kind() != targetKind {
				return errors.New(fmt.Sprintf(
					"The type of value of '%v' not match the type in the struct %v", fieldName, objectType.Name()))
			}
		} else {
			value, err := getDefaultFunctionFor(fieldName, field.Type)
			if err != nil {
				return err
			} else {
				_factoryValue.value = value
			}
		}
	}
	return nil
}

func getDefaultFunctionFor(name string, _type reflect.Type) (interface{}, error) {
	switch _type.Kind() {
	case reflect.String:
		return ProduceString, nil
	case reflect.Int:
		return ProduceInt, nil
	case reflect.Int8:
		return ProduceInt8, nil
	case reflect.Int16:
		return ProduceInt16, nil
	case reflect.Int32:
		return ProduceInt32, nil
	case reflect.Int64:
		return ProduceInt64, nil
	case reflect.Uint:
		return ProduceUint, nil
	case reflect.Uint8:
		return ProduceUint8, nil
	case reflect.Uint16:
		return ProduceUint16, nil
	case reflect.Uint32:
		return ProduceUint32, nil
	case reflect.Uint64:
		return ProduceUint64, nil
	case reflect.Float32:
		return ProduceFloat32, nil
	case reflect.Float64:
		return ProduceFloat64, nil
	case reflect.Bool:
		return ProduceBool, nil
	case reflect.Struct:
		return func(object interface{}) {
			valuePointer := reflect.New(_type).Interface()
			Factory(name).Produce(valuePointer)
			field := reflect.Indirect(reflect.ValueOf(object)).FieldByName(name)
			field.Set(reflect.ValueOf(valuePointer).Elem())
		}, nil
	case reflect.Ptr:
		switch _type.String() {
		case "*string":
			return ProduceStringPointer, nil
		case "*int":
			return ProduceIntPointer, nil
		case "*int8":
			return ProduceInt8Pointer, nil
		case "*int16":
			return ProduceInt16Pointer, nil
		case "*int32":
			return ProduceInt32Pointer, nil
		case "*int64":
			return ProduceInt64Pointer, nil
		case "*uint":
			return ProduceUintPointer, nil
		case "*uint8":
			return ProduceUint8Pointer, nil
		case "*uint16":
			return ProduceUint16Pointer, nil
		case "*uint32":
			return ProduceUint32Pointer, nil
		case "*uint64":
			return ProduceUint64Pointer, nil
		case "*float32":
			return ProduceFloat32Pointer, nil
		case "*float64":
			return ProduceFloat64Pointer, nil
		case "*bool":
			return ProduceBoolPointer, nil
		default:
			return func(object interface{}) {
				valuePointer := reflect.New(_type.Elem()).Interface()
				Factory(name).Produce(valuePointer)
				field := reflect.Indirect(reflect.ValueOf(object)).FieldByName(name)
				field.Set(reflect.ValueOf(valuePointer))
			}, nil
		}
	}

	return nil, errors.New(fmt.Sprintf(
		"There is not a default function for type %v", _type.String()))
}
