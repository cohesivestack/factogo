package factogo

type factoryValue struct {
	name   string
	value  interface{}
	params []interface{}
}

type FactoryObject struct {
	name   string
	object interface{}
	values map[string]*factoryValue
}

var factories map[string]*FactoryObject

func init() {
	Clear()
}

func Clear() {
	factories = make(map[string]*FactoryObject)
}

func Factory(name string) *FactoryObject {
	factory := &FactoryObject{name: name}
	factory.values = make(map[string]*factoryValue)
	return factory
}
