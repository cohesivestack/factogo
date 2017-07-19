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
