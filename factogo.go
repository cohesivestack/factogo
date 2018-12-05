/*
Package factogo implements a simple, flexible and useful object generator

Example:
  type User struct {
    username string
    isAdmin  bool a
  }

More Information in https://factogo.cohesivestack.com

Copyright Carlos Forero - All Rights Reserved
*/
package factogo

type factoryValue struct {
	name   string
	value  interface{}
	params []interface{}
}

type factoryInstance struct {
	notPersist           bool
	persistFunction      func(product interface{})
	afterPersistFunction func(product interface{})
	name                 string
	object               interface{}
	values               map[string]*factoryValue
}

var factories map[string]*factoryInstance
var persistFunction func(product interface{})
var afterPersistFunction func(product interface{})

func init() {
	Clear()
}

/*
Clear removes all designed Factory instances. It also unregisters the default
Persist function and the default AfterPersist callback.
*/
func Clear() {
	factories = make(map[string]*factoryInstance)
	persistFunction = nil
	afterPersistFunction = nil
}

/*
Factory get or creates a Factory Instance depending on whether it was already
designed or not.
*/
func Factory(name string) *factoryInstance {
	factory := &factoryInstance{name: name}
	factory.values = make(map[string]*factoryValue)
	return factory
}

/*
Factory get or creates a Factory Instance depending on whether it was already
designed or not.
*/
func Produce(product interface{}) {
	fi := &factoryInstance{object: product}
	fi.values = make(map[string]*factoryValue)
	_ = fi.Produce(product)
}

/*
Persist registers the default Persist function for all designed Factories.

Registering a default Persist function is optional.

A Factory can override the default Persist function.
*/
func Persist(function func(product interface{})) {
	persistFunction = function
}

/*
AfterPersist registers the default AfterPersist callback used by any Factory.

Registering a default Persist callback is optional.

A Factory can override the default AfterPersist callback.
*/
func AfterPersist(function func(product interface{})) {
	afterPersistFunction = function
}
