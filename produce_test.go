package factogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var modelA *ModelA

func TestProduce(t *testing.T) {
	Clear()
	Factory("modelA").Design(&ModelA{})
	modelA = &ModelA{}

	assert.NoError(t, Factory("modelA").Produce(modelA), "Should be produced")
}

func TestProduceTwoFactories(t *testing.T) {
	Clear()
	Factory("modelA").Design(&ModelA{})
	Factory("modelB").Design(&ModelB{})

	modelA := &ModelA{}
	modelB := &ModelB{}
	assert.NoError(t, Factory("modelA").Produce(modelA), "Should be produced")
	assert.NoError(t, Factory("modelB").Produce(modelB), "Should be produced")
}
