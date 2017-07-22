package factogo

type factoryValue struct {
	name   string
	value  interface{}
	params []interface{}
}

type FactoryInstance struct {
	notPersist           bool
	persistFunction      func(product interface{})
	afterPersistFunction func(product interface{})
	name                 string
	object               interface{}
	values               map[string]*factoryValue
}

var factories map[string]*FactoryInstance
var persistFunction func(product interface{})
var afterPersistFunction func(product interface{})

func init() {
	Clear()
}

func Clear() {
	factories = make(map[string]*FactoryInstance)
	persistFunction = nil
	afterPersistFunction = nil
}

func Factory(name string) *FactoryInstance {
	factory := &FactoryInstance{name: name}
	factory.values = make(map[string]*factoryValue)
	return factory
}

func Persist(function func(product interface{})) {
	persistFunction = function
}

func AfterPersist(function func(product interface{})) {
	afterPersistFunction = function
}
