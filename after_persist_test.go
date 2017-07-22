package factogo

import (
	"testing"

	"reflect"

	"github.com/stretchr/testify/assert"
)

type afterPersistFramework struct {
	productPersistent interface{}
}

func (this *afterPersistFramework) Persist(product interface{}) {
	this.productPersistent = product
}

type afterPersistFrameworkEmpty struct {
}

func (this *afterPersistFrameworkEmpty) Persist(product interface{}) {
	// leave empty with test purposes
}

func TestAfterPersistGlobally(t *testing.T) {
	Clear()

	messageToTest := "Has entered in after persist"
	persistence := &afterPersistFramework{}

	Persist(persistence.Persist)
	AfterPersist(func(product interface{}) {
		if reflect.ValueOf(product).Type() == reflect.TypeOf(&ModelA{}) {
			_modelA := product.(*ModelA)
			_modelA.String = messageToTest
		}
	})
	modelA = &ModelA{}

	assert.NoError(t, Factory("modelA").Design(modelA), "Should be designed")
	assert.NoError(t, Factory("modelA").Produce(modelA), "Should be produced")

	assert.Equal(t, messageToTest, modelA.String, "Should be equal")
}

func TestAfterPersistOnDesign(t *testing.T) {
	messageToTest := "Has entered in after persist"
	persistence := &afterPersistFramework{}

	Persist(persistence.Persist)
	modelA = &ModelA{}

	_afterPersist := func(product interface{}) {
		if reflect.ValueOf(product).Type() == reflect.TypeOf(&ModelA{}) {
			_modelA := product.(*ModelA)
			_modelA.String = messageToTest
		}
	}

	assert.NoError(t, Factory("modelA").AfterPersist(_afterPersist).Design(modelA),
		"Should be designed")
	assert.NoError(t, Factory("modelA").Produce(modelA), "Should be produced")

	assert.Equal(t, messageToTest, modelA.String, "Should be equal")
}

func TestAfterPersistOnProduce(t *testing.T) {
	messageToTest := "Has entered in after persist"
	persistence := &afterPersistFramework{}

	Persist(persistence.Persist)
	modelA = &ModelA{}

	_afterPersist := func(product interface{}) {
		if reflect.ValueOf(product).Type() == reflect.TypeOf(&ModelA{}) {
			_modelA := product.(*ModelA)
			_modelA.String = messageToTest
		}
	}

	assert.NoError(t, Factory("modelA").Design(modelA),
		"Should be designed")
	assert.NoError(t, Factory("modelA").AfterPersist(_afterPersist).Produce(modelA),
		"Should be produced")

	assert.Equal(t, messageToTest, modelA.String, "Should be equal")
}

func TestAfterPersistOverridingGlobal(t *testing.T) {
	messageToTest := "Has entered in after persist"
	persistence := &afterPersistFramework{}

	Persist(persistence.Persist)
	modelA = &ModelA{}

	AfterPersist(func(product interface{}) {
		if reflect.ValueOf(product).Type() == reflect.TypeOf(&ModelA{}) {
			_modelA := product.(*ModelA)
			_modelA.String = "This value never should be assigned"
		}
	})

	_afterPersist := func(product interface{}) {
		if reflect.ValueOf(product).Type() == reflect.TypeOf(&ModelA{}) {
			_modelA := product.(*ModelA)
			_modelA.String = messageToTest
		}
	}

	assert.NoError(t, Factory("modelA").AfterPersist(_afterPersist).Design(modelA),
		"Should be designed")
	assert.NoError(t, Factory("modelA").Produce(modelA), "Should be produced")

	assert.Equal(t, messageToTest, modelA.String, "Should be equal")
}

func TestAfterPersistOverridingDesign(t *testing.T) {
	messageToTest := "Has entered in after persist"
	persistence := &afterPersistFramework{}

	Persist(persistence.Persist)
	modelA = &ModelA{}

	_afterPersistOnDesign := func(product interface{}) {
		if reflect.ValueOf(product).Type() == reflect.TypeOf(&ModelA{}) {
			_modelA := product.(*ModelA)
			_modelA.String = "This value never should be assigned"
		}
	}

	_afterPersist := func(product interface{}) {
		if reflect.ValueOf(product).Type() == reflect.TypeOf(&ModelA{}) {
			_modelA := product.(*ModelA)
			_modelA.String = messageToTest
		}
	}

	assert.NoError(t, Factory("modelA").AfterPersist(_afterPersistOnDesign).Design(modelA),
		"Should be designed")
	assert.NoError(t, Factory("modelA").AfterPersist(_afterPersist).Produce(modelA),
		"Should be produced")

	assert.Equal(t, messageToTest, modelA.String, "Should be equal")
}
