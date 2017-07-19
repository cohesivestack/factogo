package factogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduceFunction(t *testing.T) {
	Clear()
	modelA = &ModelA{}

	intValue := 10
	int8Value := int8(intValue)
	int16Value := int16(intValue)
	int32Value := int32(intValue)
	int64Value := int64(intValue)
	uintValue := uint(intValue)
	uint8Value := uint8(intValue)
	uint16Value := uint16(intValue)
	uint32Value := uint32(intValue)
	uint64Value := uint64(intValue)
	float32Value := float32(10.50)
	float64Value := float64(float32Value)
	stringValue := "Hello"
	boolValue := true

	modelBValue := ModelB{}
	modelCPointer := &ModelC{}

	// Default function
	Factory("modelA").
		Set("String", func() string { return stringValue }).
		Set("Int", func() int { return intValue }).
		Set("Int8", func() int8 { return int8Value }).
		Set("Int16", func() int16 { return int16Value }).
		Set("Int32", func() int32 { return int32Value }).
		Set("Int64", func() int64 { return int64Value }).
		Set("Uint", func() uint { return uintValue }).
		Set("Uint8", func() uint8 { return uint8Value }).
		Set("Uint16", func() uint16 { return uint16Value }).
		Set("Uint32", func() uint32 { return uint32Value }).
		Set("Uint64", func() uint64 { return uint64Value }).
		Set("Float32", func() float32 { return float32Value }).
		Set("Float64", func() float64 { return float64Value }).
		Set("ModelB", func() ModelB { return modelBValue }).
		Set("Bool", func() bool { return boolValue }).
		Set("StringPointer", func() *string { return &stringValue }).
		Set("IntPointer", func() *int { return &intValue }).
		Set("Int8Pointer", func() *int8 { return &int8Value }).
		Set("Int16Pointer", func() *int16 { return &int16Value }).
		Set("Int32Pointer", func() *int32 { return &int32Value }).
		Set("Int64Pointer", func() *int64 { return &int64Value }).
		Set("UintPointer", func() *uint { return &uintValue }).
		Set("Uint8Pointer", func() *uint8 { return &uint8Value }).
		Set("Uint16Pointer", func() *uint16 { return &uint16Value }).
		Set("Uint32Pointer", func() *uint32 { return &uint32Value }).
		Set("Uint64Pointer", func() *uint64 { return &uint64Value }).
		Set("Float32Pointer", func() *float32 { return &float32Value }).
		Set("Float64Pointer", func() *float64 { return &float64Value }).
		Set("BoolPointer", func() *bool { return &boolValue }).
		Set("ModelC", func() *ModelC { return modelCPointer }).
		Design(&ModelA{})

	assert.NoError(t, Factory("modelA").Produce(modelA), "Should be produced")
	assert.Equal(t, stringValue, modelA.String, "Should be equal")
	assert.Equal(t, intValue, modelA.Int, "Should be equal")
	assert.Equal(t, int8Value, modelA.Int8, "Should be equal")
	assert.Equal(t, int16Value, modelA.Int16, "Should be equal")
	assert.Equal(t, int32Value, modelA.Int32, "Should be equal")
	assert.Equal(t, int64Value, modelA.Int64, "Should be equal")
	assert.Equal(t, uintValue, modelA.Uint, "Should be equal")
	assert.Equal(t, uint8Value, modelA.Uint8, "Should be equal")
	assert.Equal(t, uint16Value, modelA.Uint16, "Should be equal")
	assert.Equal(t, uint32Value, modelA.Uint32, "Should be equal")
	assert.Equal(t, uint64Value, modelA.Uint64, "Should be equal")
	assert.Equal(t, float32Value, modelA.Float32, "Should be equal")
	assert.Equal(t, float64Value, modelA.Float64, "Should be equal")
	assert.Equal(t, boolValue, modelA.Bool, "Should be equal")
	assert.Equal(t, modelBValue, modelA.ModelB, "Should be equal")

	assert.Equal(t, stringValue, *modelA.StringPointer, "Should be equal")
	assert.Equal(t, intValue, *modelA.IntPointer, "Should be equal")
	assert.Equal(t, int8Value, *modelA.Int8Pointer, "Should be equal")
	assert.Equal(t, int16Value, *modelA.Int16Pointer, "Should be equal")
	assert.Equal(t, int32Value, *modelA.Int32Pointer, "Should be equal")
	assert.Equal(t, int64Value, *modelA.Int64Pointer, "Should be equal")
	assert.Equal(t, uintValue, *modelA.UintPointer, "Should be equal")
	assert.Equal(t, uint8Value, *modelA.Uint8Pointer, "Should be equal")
	assert.Equal(t, uint16Value, *modelA.Uint16Pointer, "Should be equal")
	assert.Equal(t, uint32Value, *modelA.Uint32Pointer, "Should be equal")
	assert.Equal(t, uint64Value, *modelA.Uint64Pointer, "Should be equal")
	assert.Equal(t, float32Value, *modelA.Float32Pointer, "Should be equal")
	assert.Equal(t, float64Value, *modelA.Float64Pointer, "Should be equal")
	assert.Equal(t, boolValue, *modelA.BoolPointer, "Should be equal")
	assert.Equal(t, modelCPointer, modelA.ModelC, "Should be equal")
}
