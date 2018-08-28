package factogo

/*
Persist registers the function, over a designed Factory, used to persist a
product Struct.

Note that Persist won't call the function to persist, it only registers
the function. The function is called when the Produce() method is called over
a designed Factory.

In case Factogo has registered a default Persist function, registering a Persist
function, over the Factory instance, overrides the default Persist function.
*/
func (fi *factoryInstance) Persist(function func(product interface{})) *factoryInstance {
	fi.notPersist = false
	fi.persistFunction = function
	return fi
}

/*
NotPersist avoids to call a Persist function, even when there is a registered
default Persist function.
*/
func (fi *factoryInstance) NotPersist() *factoryInstance {
	fi.notPersist = true
	fi.persistFunction = nil
	return fi
}
