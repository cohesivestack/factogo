package factogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuto(t *testing.T) {
	Clear()

	Factory("modelD").Auto().Design(&ModelD{})

	modelD := &ModelD{}
	assert.NoError(t, Factory("modelD").Produce(modelD), "Should be produced")

	assert.NotZero(t, modelD.Int)
	assert.NotEmpty(t, modelD.String)
}

func TestAutoAsDefault(t *testing.T) {
	Clear()

	Factory("modelD").Design(&ModelD{})

	modelD := &ModelD{}
	assert.NoError(t, Factory("modelD").Produce(modelD), "Should be produced")

	assert.NotZero(t, modelD.Int)
	assert.NotEmpty(t, modelD.String)
}

func TestAutoWithAnonymous(t *testing.T) {
	Clear()

	modelD := &ModelD{}
	assert.NoError(t, Auto().Produce(modelD), "Should be produced")

	assert.NotZero(t, modelD.Int)
	assert.NotEmpty(t, modelD.String)
}

func TestAutoWithSet(t *testing.T) {
	Clear()

	Factory("modelD").Auto().Set("Int", 123456789).Design(&ModelD{})

	modelD := &ModelD{}
	assert.NoError(t, Factory("modelD").Produce(modelD), "Should be produced")

	assert.Equal(t, 123456789, modelD.Int)
	assert.NotEmpty(t, modelD.String)
}

func TestAutoWithSetWithAnonymous(t *testing.T) {
	Clear()

	modelD := &ModelD{}
	assert.NoError(t, Auto().Set("Int", 123456789).Produce(modelD), "Should be produced")

	assert.Equal(t, 123456789, modelD.Int)
	assert.NotEmpty(t, modelD.String)
}
