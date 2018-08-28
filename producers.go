package factogo

import (
	"math/rand"
	"time"
)

const defaultCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func seededRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func ProduceString() string {
	length := 20

	chars := make([]byte, length)
	for i := range chars {
		chars[i] = defaultCharset[seededRand().Intn(len(defaultCharset))]
	}
	return string(chars)
}

func ProduceInt() int {
	return seededRand().Int()
}

func ProduceInt8() int8 {
	return int8(seededRand().Intn(127))
}

func ProduceInt16() int16 {
	return int16(seededRand().Intn(32767))
}

func ProduceInt32() int32 {
	return seededRand().Int31()
}

func ProduceInt64() int64 {
	return seededRand().Int63()
}

func ProduceUint() uint {
	return uint(seededRand().Int())
}

func ProduceUint8() uint8 {
	return uint8(seededRand().Intn(255))
}

func ProduceUint16() uint16 {
	return uint16(seededRand().Intn(65535))
}

func ProduceUint32() uint32 {
	return seededRand().Uint32()
}

func ProduceUint64() uint64 {
	return seededRand().Uint64()
}

func ProduceFloat32() float32 {
	return seededRand().Float32()
}

func ProduceFloat64() float64 {
	return seededRand().Float64()
}

func ProduceBool() bool {
	var val bool
	if seededRand().Intn(1) == 1 {
		val = true
	}
	return val
}

func ProduceStringPointer() *string {
	val := ProduceString()
	return &val
}

func ProduceIntPointer() *int {
	val := ProduceInt()
	return &val
}

func ProduceInt8Pointer() *int8 {
	val := ProduceInt8()
	return &val
}

func ProduceInt16Pointer() *int16 {
	val := ProduceInt16()
	return &val
}

func ProduceInt32Pointer() *int32 {
	val := ProduceInt32()
	return &val
}

func ProduceInt64Pointer() *int64 {
	val := ProduceInt64()
	return &val
}

func ProduceUintPointer() *uint {
	val := ProduceUint()
	return &val
}

func ProduceUint8Pointer() *uint8 {
	val := ProduceUint8()
	return &val
}

func ProduceUint16Pointer() *uint16 {
	val := ProduceUint16()
	return &val
}

func ProduceUint32Pointer() *uint32 {
	val := ProduceUint32()
	return &val
}

func ProduceUint64Pointer() *uint64 {
	val := ProduceUint64()
	return &val
}

func ProduceFloat32Pointer() *float32 {
	val := ProduceFloat32()
	return &val
}

func ProduceFloat64Pointer() *float64 {
	val := ProduceFloat64()
	return &val
}

func ProduceBoolPointer() *bool {
	val := ProduceBool()
	return &val
}
