package factogo

func (this *FactoryInstance) Set(
	name string, params ...interface{}) *FactoryInstance {

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

	this.values[name] = &factoryValue{name: name, value: value, params: params}
	return this
}
