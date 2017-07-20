package factogo

import (
	"errors"
	"fmt"
	"reflect"
)

func (this *FactoryInstance) Design(object interface{}) error {
	err := checkValues(object, this)
	if err != nil {
		return err
	}
	factories[this.name] = this
	return nil
}

func checkValues(object interface{}, factory *FactoryInstance) error {

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
		return DefaultStringFunction, nil
	case reflect.Int:
		return DefaultIntFunction, nil
	case reflect.Int8:
		return DefaultInt8Function, nil
	case reflect.Int16:
		return DefaultInt16Function, nil
	case reflect.Int32:
		return DefaultInt32Function, nil
	case reflect.Int64:
		return DefaultInt64Function, nil
	case reflect.Uint:
		return DefaultUintFunction, nil
	case reflect.Uint8:
		return DefaultUint8Function, nil
	case reflect.Uint16:
		return DefaultUint16Function, nil
	case reflect.Uint32:
		return DefaultUint32Function, nil
	case reflect.Uint64:
		return DefaultUint64Function, nil
	case reflect.Float32:
		return DefaultFloat32Function, nil
	case reflect.Float64:
		return DefaultFloat64Function, nil
	case reflect.Bool:
		return DefaultBoolFunction, nil
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
			return DefaultStringPointerFunction, nil
		case "*int":
			return DefaultIntPointerFunction, nil
		case "*int8":
			return DefaultInt8PointerFunction, nil
		case "*int16":
			return DefaultInt16PointerFunction, nil
		case "*int32":
			return DefaultInt32PointerFunction, nil
		case "*int64":
			return DefaultInt64PointerFunction, nil
		case "*uint":
			return DefaultUintPointerFunction, nil
		case "*uint8":
			return DefaultUint8PointerFunction, nil
		case "*uint16":
			return DefaultUint16PointerFunction, nil
		case "*uint32":
			return DefaultUint32PointerFunction, nil
		case "*uint64":
			return DefaultUint64PointerFunction, nil
		case "*float32":
			return DefaultFloat32PointerFunction, nil
		case "*float64":
			return DefaultFloat64PointerFunction, nil
		case "*bool":
			return DefaultBoolPointerFunction, nil
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
