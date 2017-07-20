package factogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type persistenceFramework struct {
	productPersistent interface{}
}

func (this *persistenceFramework) Persist(product interface{}) {
	this.productPersistent = product
}

type persistenceFrameworkEmpty struct {
}

func (this *persistenceFrameworkEmpty) Persist(product interface{}) {
	// leave empty with test purposes
}

func TestPersistGlobally(t *testing.T) {
	Clear()

	persistence := &persistenceFramework{}

	Persist(persistence.Persist)
	modelA = &ModelA{}

	assert.Nil(t, persistence.productPersistent,
		"the productPersistent should not exist")

	assert.NoError(t, Factory("modelA").Design(modelA), "Should be designed")
	assert.NoError(t, Factory("modelA").Produce(modelA), "Should be produced")

	assert.NotNil(t, persistence.productPersistent,
		"the productPersistent should exist")

	assert.IsType(t, modelA, persistence.productPersistent,
		"modelA should be persisted")
}

func TestPersistOnDesign(t *testing.T) {
	Clear()

	persistence := &persistenceFramework{}

	Persist(nil)
	assert.Nil(t, persistFunction, "the global persistFunction should be nil")
	modelA = &ModelA{}

	assert.Nil(t, persistence.productPersistent,
		"the productPersistent should not exist")

	assert.NoError(t, Factory("modelA").Persist(persistence.Persist).Design(modelA),
		"Should be designed")
	assert.NoError(t, Factory("modelA").Produce(modelA), "Should be produced")

	assert.NotNil(t, persistence.productPersistent,
		"the productPersistent should exist")

	assert.IsType(t, modelA, persistence.productPersistent,
		"modelA should be persisted")
}

func TestPersistOnProduce(t *testing.T) {
	Clear()

	persistence := &persistenceFramework{}

	Persist(nil)
	assert.Nil(t, persistFunction, "the global persistFunction should be nil")
	modelA = &ModelA{}

	assert.Nil(t, persistence.productPersistent,
		"the productPersistent should not exist")

	assert.NoError(t, Factory("modelA").Design(modelA),
		"Should be designed")
	assert.NoError(t, Factory("modelA").Persist(persistence.Persist).Produce(modelA), "Should be produced")

	assert.NotNil(t, persistence.productPersistent,
		"the productPersistent should exist")

	assert.IsType(t, modelA, persistence.productPersistent,
		"modelA should be persisted")
}

func TestPersistOverridingGlobal(t *testing.T) {
	Clear()

	persistence := &persistenceFramework{}
	persistenceEmpty := &persistenceFrameworkEmpty{}

	Persist(persistenceEmpty.Persist)
	assert.NotNil(t, persistFunction, "the global persistFunction should be not nil")
	modelA = &ModelA{}

	assert.Nil(t, persistence.productPersistent,
		"the productPersistent should not exist")

	assert.NoError(t, Factory("modelA").Persist(persistence.Persist).Design(modelA),
		"Should be designed")
	assert.NoError(t, Factory("modelA").Produce(modelA), "Should be produced")

	assert.NotNil(t, persistence.productPersistent,
		"the productPersistent should exist")

	assert.IsType(t, modelA, persistence.productPersistent,
		"modelA should be persisted")
}

func TestPersistOverridingDesign(t *testing.T) {
	Clear()

	persistence := &persistenceFramework{}
	persistenceEmpty := &persistenceFrameworkEmpty{}

	Persist(nil)
	assert.Nil(t, persistFunction, "the global persistFunction should be nil")
	modelA = &ModelA{}

	assert.Nil(t, persistence.productPersistent,
		"the productPersistent should not exist")

	assert.NoError(t, Factory("modelA").Persist(persistenceEmpty.Persist).Design(modelA),
		"Should be designed")
	assert.NoError(t, Factory("modelA").Persist(persistence.Persist).Produce(modelA), "Should be produced")

	assert.NotNil(t, persistence.productPersistent,
		"the productPersistent should exist")

	assert.IsType(t, modelA, persistence.productPersistent,
		"modelA should be persisted")
}

func TestNotPersistOnDesign(t *testing.T) {
	Clear()

	persistence := &persistenceFramework{}

	Persist(persistence.Persist)
	assert.NotNil(t, persistFunction, "the global persistFunction should be not nil")
	modelA = &ModelA{}

	assert.Nil(t, persistence.productPersistent,
		"the productPersistent should not exist")

	assert.NoError(t, Factory("modelA").NotPersist().Design(modelA),
		"Should be designed")
	assert.NoError(t, Factory("modelA").Produce(modelA), "Should be produced")

	assert.Nil(t, persistence.productPersistent,
		"the productPersistent should not exist")
}

func TestNotPersistOnProduce(t *testing.T) {
	Clear()

	persistence := &persistenceFramework{}

	Persist(persistence.Persist)
	assert.NotNil(t, persistFunction, "the global persistFunction should be not nil")
	modelA = &ModelA{}

	assert.Nil(t, persistence.productPersistent,
		"the productPersistent should not exist")

	assert.NoError(t, Factory("modelA").Design(modelA),
		"Should be designed")
	assert.NoError(t, Factory("modelA").NotPersist().Produce(modelA), "Should be produced")

	assert.Nil(t, persistence.productPersistent,
		"the productPersistent should not exist")
}

func TestChangeFromNotPersistToPersist(t *testing.T) {
	Clear()

	persistence := &persistenceFramework{}

	Persist(persistence.Persist)
	assert.NotNil(t, persistFunction, "the global persistFunction should be not nil")
	modelA = &ModelA{}

	assert.Nil(t, persistence.productPersistent,
		"the productPersistent should not exist")

	assert.NoError(t, Factory("modelA").NotPersist().Design(modelA),
		"Should be designed")
	assert.NoError(t, Factory("modelA").Persist(persistence.Persist).Produce(modelA),
		"Should be produced")

	assert.NotNil(t, persistence.productPersistent,
		"the productPersistent should exist")

	assert.IsType(t, modelA, persistence.productPersistent,
		"modelA should be persisted")
}
