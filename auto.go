package factogo

func (fi *factoryInstance) Auto() *factoryInstance {
	fi.isAuto = true
	return fi
}

func Auto() *factoryInstance {
	return factoryAnonymous().Auto()
}
