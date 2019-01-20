package factogo

import (
	"fmt"
	"reflect"
)

/*
Creates a new product struct from an anonymous factory design
*/
func Produce(product interface{}) error {
	fi := &factoryInstance{object: product, isAnonymous: true}
	fi.values = make(map[string]*factoryValue)
	return fi.Produce(product)
}

/*
Produce invokes a designed Factory to create a new product Struct.

The values for the product Struct are copied in the target parameter passed
to the Produce function.

Example:
	Factory("staff").Design(Staff{})
	target := &Staff{}
	Factory("staff").Produce(target)
*/
func (fi *factoryInstance) Produce(product interface{}) error {
	modifiedFactory := fi

	if fi.isAnonymous {
		err := fi.Design(product)
		if err != nil {
			return err
		}
	} else {
		if _, ok := factories[fi.name]; !ok {
			return fmt.Errorf("'%s' designed Factory doesn't exist", fi.name)
		}

		originalFactory := factories[fi.name]

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
	}

	for _, value := range modifiedFactory.values {
		valueType := reflect.TypeOf(value.value)
		field := reflect.Indirect(reflect.ValueOf(product)).FieldByName(value.name)

		if valueType.Kind() == reflect.Func {

			if reflect.TypeOf(value.value).NumOut() > 0 {
				_value := reflect.ValueOf(value.value).Call([]reflect.Value{})[0].Interface()
				field.Set(reflect.ValueOf(_value))
			} else {
				reflect.ValueOf(value.value).Call([]reflect.Value{reflect.ValueOf(product)})
			}
		} else {
			field.Set(reflect.ValueOf(value.value))
		}
	}

	if !fi.notPersist && (fi.persistFunction != nil || persistFunction != nil) {
		if fi.persistFunction == nil {
			fi.persistFunction = persistFunction
		}
		fi.persistFunction(product)

		if fi.afterPersistFunction == nil {
			fi.afterPersistFunction = afterPersistFunction
		}
		if fi.afterPersistFunction != nil {
			fi.afterPersistFunction(product)
		}
	}

	return nil
}
