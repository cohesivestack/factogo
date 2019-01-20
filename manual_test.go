package factogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManual(t *testing.T) {
	Clear()

	Factory("modelD").Manual().Design(&ModelD{})

	modelD := &ModelD{}
	assert.NoError(t, Factory("modelD").Produce(modelD), "Should be produced")

	assert.Zero(t, modelD.Int)
	assert.Empty(t, modelD.String)
}

func TestManualWithSet(t *testing.T) {
	Clear()

	Factory("modelD").Manual().Set("Int", 123).Design(&ModelD{})

	modelD := &ModelD{}
	assert.NoError(t, Factory("modelD").Produce(modelD), "Should be produced")

	assert.Equal(t, 123, modelD.Int)
	assert.Empty(t, modelD.String)
}
