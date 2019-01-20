package factogo

/*
Set specifies a Struct's field which should be assigned by a Factory when the
Produce() method is called.

The Set() method receives two parameters:

The first parameter "name" is used to specify which is the name of the Struct
field.

The second parameter is optional and can be used to pass a value different to
the generated values produced by factogo. If the second parameter is a function
then fi function is called to get the value to assign to the field.
*/
func (fi *factoryInstance) Set(
	name string, params ...interface{}) *factoryInstance {

	var value interface{}

	if len(params) > 0 {
		value = params[0]
		if len(params) > 1 {
			params = params[1:]
		} else {
			params = make([]interface{}, 0)
		}
	} else {
		value = nil
	}

	fi.values[name] = &factoryValue{name: name, value: value, params: params}
	return fi
}

/*
Set specifies a Struct's field which should be assigned by a anonymous Factory when the
Produce() method is called.

The Set() method receives two parameters:

The first parameter "name" is used to specify which is the name of the Struct
field.

The second parameter is optional and can be used to pass a value different to
the generated values produced by factogo. If the second parameter is a function
then fi function is called to get the value to assign to the field.
*/

func Set(name string, params ...interface{}) *factoryInstance {
	return factoryAnonymous().Set(name, params)
}
