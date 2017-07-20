package factogo

func (this *FactoryInstance) Persist(function func(product interface{})) *FactoryInstance {
	this.notPersist = false
	this.persistFunction = function
	return this
}

func (this *FactoryInstance) NotPersist() *FactoryInstance {
	this.notPersist = true
	this.persistFunction = nil
	return this
}
