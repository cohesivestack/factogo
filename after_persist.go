package factogo

/*
AfterPersist registers a callback, which is called after the Persist function
is executed.

AfterPersist receives the product Struct which should have been duly persisted
using the Persist function.

If any Persist function is registered or NotPersist() was called over the
designed Factory, then the AfterPersist callback won't be called.
*/
func (this *factoryInstance) AfterPersist(function func(product interface{})) *factoryInstance {
	this.afterPersistFunction = function
	return this
}
