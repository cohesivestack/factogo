package factogo

func (fi *factoryInstance) Manual() *factoryInstance {
	fi.isAuto = false
	return fi
}

func Manual() *factoryInstance {
	return factoryAnonymous().Manual()
}
