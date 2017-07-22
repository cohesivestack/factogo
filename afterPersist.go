package factogo

func (this *FactoryInstance) AfterPersist(function func(product interface{})) *FactoryInstance {
	this.afterPersistFunction = function
	return this
}
