package factogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduceDefault(t *testing.T) {
	Clear()
	modelA = &ModelA{}

	Factory("ModelB").
		Set("String").
		Design(&ModelB{})

	Factory("ModelC").
		Set("String").
		Design(&ModelC{})

	Factory("ModelA").
		Set("String").
		Set("Int").
		Set("Int8").
		Set("Int16").
		Set("Int32").
		Set("Int64").
		Set("Uint").
		Set("Uint8").
		Set("Uint16").
		Set("Uint32").
		Set("Uint64").
		Set("Float32").
		Set("Float64").
		Set("Bool").
		Set("Time").
		Set("ModelB").
		Set("IntArray").
		Set("ModelBArray").
		Set("StringMap").
		Set("IntPointerArray").
		Set("ModelBPointerArray").
		Set("StringPointerMap").
		Set("StringPointer").
		Set("IntPointer").
		Set("Int8Pointer").
		Set("Int16Pointer").
		Set("Int32Pointer").
		Set("Int64Pointer").
		Set("UintPointer").
		Set("Uint8Pointer").
		Set("Uint16Pointer").
		Set("Uint32Pointer").
		Set("Uint64Pointer").
		Set("Float32Pointer").
		Set("Float64Pointer").
		Set("BoolPointer").
		Set("ModelC").
		Set("IntArrayPointer").
		Set("ModelCArrayPointer").
		Set("StringMapPointer").
		Set("IntPointerArrayPointer").
		Set("ModelCPointerArrayPointer").
		Set("StringPointerMapPointer").
		Design(&ModelA{})

	assert.NoError(t, Factory("ModelA").Produce(modelA), "Should be produced")
	assert.NotEmpty(t, modelA.String, "Should be not empty")
	assert.NotEmpty(t, modelA.Int, "Should be not empty")
	assert.NotEmpty(t, modelA.Int8, "Should be not empty")
	assert.NotEmpty(t, modelA.Int16, "Should be not empty")
	assert.NotEmpty(t, modelA.Int32, "Should be not empty")
	assert.NotEmpty(t, modelA.Int64, "Should be not empty")
	assert.NotEmpty(t, modelA.Uint, "Should be not empty")
	assert.NotEmpty(t, modelA.Uint8, "Should be not empty")
	assert.NotEmpty(t, modelA.Uint16, "Should be not empty")
	assert.NotEmpty(t, modelA.Uint32, "Should be not empty")
	assert.NotEmpty(t, modelA.Uint64, "Should be not empty")
	assert.NotEmpty(t, modelA.Float32, "Should be not empty")
	assert.NotEmpty(t, modelA.Float64, "Should be not empty")
	assert.NotNil(t, modelA.ModelB, "Should be not empty")
	assert.NotZero(t, modelA.Time, "Should be not empty")

	assert.Len(t, modelA.IntArray, 0, "Should be empty")
	assert.Len(t, modelA.ModelBArray, 0, "Should be empty")
	assert.Len(t, modelA.StringMap, 0, "Should be empty")
	assert.Len(t, modelA.IntPointerArray, 0, "Should be empty")
	assert.Len(t, modelA.ModelBPointerArray, 0, "Should be empty")
	assert.Len(t, modelA.StringPointerMap, 0, "Should be empty")

	assert.NotEmpty(t, modelA.ModelB.String, "Should be not empty")
	assert.NotEmpty(t, *modelA.StringPointer, "Should be not empty")
	assert.NotEmpty(t, *modelA.IntPointer, "Should be not empty")
	assert.NotEmpty(t, *modelA.Int8Pointer, "Should be not empty")
	assert.NotEmpty(t, *modelA.Int16Pointer, "Should be not empty")
	assert.NotEmpty(t, *modelA.Int32Pointer, "Should be not empty")
	assert.NotEmpty(t, *modelA.Int64Pointer, "Should be not empty")
	assert.NotEmpty(t, *modelA.UintPointer, "Should be not empty")
	assert.NotEmpty(t, *modelA.Uint8Pointer, "Should be not empty")
	assert.NotEmpty(t, *modelA.Uint16Pointer, "Should be not empty")
	assert.NotEmpty(t, *modelA.Uint32Pointer, "Should be not empty")
	assert.NotEmpty(t, *modelA.Uint64Pointer, "Should be not empty")
	assert.NotEmpty(t, *modelA.Float32Pointer, "Should be not empty")
	assert.NotEmpty(t, *modelA.Float64Pointer, "Should be not empty")

	assert.NotNil(t, modelA.BoolPointer, "Should be not empty")
	assert.NotZero(t, *modelA.TimePointer, "Should be not empty")

	assert.NotNil(t, modelA.IntArrayPointer, "Should be not empty")
	assert.NotNil(t, modelA.ModelCArrayPointer, "Should be not empty")
	assert.NotNil(t, modelA.StringMapPointer, "Should be not empty")
	assert.NotNil(t, modelA.IntPointerArrayPointer, "Should be not empty")
	assert.NotNil(t, modelA.ModelCPointerArrayPointer, "Should be not empty")
	assert.NotNil(t, modelA.StringPointerMapPointer, "Should be not empty")
}
