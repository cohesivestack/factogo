package factogo

type factoryValue struct {
	name   string
	value  interface{}
	params []interface{}
}

type FactoryInstance struct {
	name   string
	object interface{}
	values map[string]*factoryValue
}

var factories map[string]*FactoryInstance

func init() {
	Clear()
}

func Clear() {
	factories = make(map[string]*FactoryInstance)
}

func Factory(name string) *FactoryInstance {
	factory := &FactoryInstance{name: name}
	factory.values = make(map[string]*factoryValue)
	return factory
}
