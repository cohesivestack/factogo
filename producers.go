package factogo

import (
	"math/rand"
	"time"
)

const defaultCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func seededRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

/*
ProduceString produces a ramdom string.
*/
func ProduceString() string {
	length := 20

	chars := make([]byte, length)
	for i := range chars {
		chars[i] = defaultCharset[seededRand().Intn(len(defaultCharset))]
	}
	return string(chars)
}

/*
ProduceInt produces a ramdom int number.
*/
func ProduceInt() int {
	return seededRand().Int()
}

/*
ProduceInt8 produces a ramdom int8 number.
*/
func ProduceInt8() int8 {
	return int8(seededRand().Intn(127))
}

/*
ProduceInt16 produces a ramdom int16 number.
*/
func ProduceInt16() int16 {
	return int16(seededRand().Intn(32767))
}

/*
ProduceInt32 produces a ramdom int32 number.
*/
func ProduceInt32() int32 {
	return seededRand().Int31()
}

/*
ProduceInt64 produces a ramdom int64 number.
*/
func ProduceInt64() int64 {
	return seededRand().Int63()
}

/*
ProduceUint produces a ramdom uint number.
*/
func ProduceUint() uint {
	return uint(seededRand().Int())
}

/*
ProduceUint8 produces a ramdom uint8 number.
*/
func ProduceUint8() uint8 {
	return uint8(seededRand().Intn(255))
}

/*
ProduceUint16 produces a ramdom uint16 number.
*/
func ProduceUint16() uint16 {
	return uint16(seededRand().Intn(65535))
}

/*
ProduceUint32 produces a ramdom uint32 number.
*/
func ProduceUint32() uint32 {
	return seededRand().Uint32()
}

/*
ProduceUint64 produces a ramdom uint64 number.
*/

func ProduceUint64() uint64 {
	return seededRand().Uint64()
}

/*
ProduceFloat32 produces a ramdom float32 number.
*/
func ProduceFloat32() float32 {
	return seededRand().Float32()
}

/*
ProduceFloat64 produces a ramdom float64 number.
*/
func ProduceFloat64() float64 {
	return seededRand().Float64()
}

/*
ProduceBool produces a ramdom boolean.
*/
func ProduceBool() bool {
	var val bool
	if seededRand().Intn(1) == 1 {
		val = true
	}
	return val
}

/*
ProduceStringPointer produces a pointer to a ramdom string.
*/
func ProduceStringPointer() *string {
	val := ProduceString()
	return &val
}

/*
ProduceIntPointer produces a pointer to a ramdom int.
*/
func ProduceIntPointer() *int {
	val := ProduceInt()
	return &val
}

/*
ProduceInt8Pointer produces a pointer to a ramdom int8.
*/
func ProduceInt8Pointer() *int8 {
	val := ProduceInt8()
	return &val
}

/*
ProduceInt16Pointer produces a pointer to a ramdom int16.
*/
func ProduceInt16Pointer() *int16 {
	val := ProduceInt16()
	return &val
}

/*
ProduceInt32Pointer produces a pointer to a ramdom int32.
*/
func ProduceInt32Pointer() *int32 {
	val := ProduceInt32()
	return &val
}

/*
ProduceInt54Pointer produces a pointer to a ramdom int64.
*/
func ProduceInt64Pointer() *int64 {
	val := ProduceInt64()
	return &val
}

/*
ProduceUintPointer produces a pointer to a ramdom uint.
*/
func ProduceUintPointer() *uint {
	val := ProduceUint()
	return &val
}

/*
ProduceUint8Pointer produces a pointer to a ramdom uint8.
*/
func ProduceUint8Pointer() *uint8 {
	val := ProduceUint8()
	return &val
}

/*
ProduceUint16Pointer produces a pointer to a ramdom uint16.
*/
func ProduceUint16Pointer() *uint16 {
	val := ProduceUint16()
	return &val
}

/*
ProduceUint32Pointer produces a pointer to a ramdom uint32.
*/
func ProduceUint32Pointer() *uint32 {
	val := ProduceUint32()
	return &val
}

/*
ProduceUint64Pointer produces a pointer to a ramdom uint64.
*/
func ProduceUint64Pointer() *uint64 {
	val := ProduceUint64()
	return &val
}

/*
ProduceFloat32Pointer produces a pointer to a ramdom float32.
*/
func ProduceFloat32Pointer() *float32 {
	val := ProduceFloat32()
	return &val
}

/*
ProduceFloat54Pointer produces a pointer to a ramdom float64.
*/
func ProduceFloat64Pointer() *float64 {
	val := ProduceFloat64()
	return &val
}

/*
ProduceBoolPointer produces a pointer to a ramdom bool.
*/
func ProduceBoolPointer() *bool {
	val := ProduceBool()
	return &val
}
