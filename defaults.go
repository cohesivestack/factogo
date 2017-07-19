package factogo

import (
	"math/rand"
	"time"
)

const defaultCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func seededRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func DefaultStringFunction() string {
	length := 20

	chars := make([]byte, length)
	for i := range chars {
		chars[i] = defaultCharset[seededRand().Intn(len(defaultCharset))]
	}
	return string(chars)
}

func DefaultIntFunction() int {
	return seededRand().Int()
}

func DefaultInt8Function() int8 {
	return int8(seededRand().Intn(127))
}

func DefaultInt16Function() int16 {
	return int16(seededRand().Intn(32767))
}

func DefaultInt32Function() int32 {
	return seededRand().Int31()
}

func DefaultInt64Function() int64 {
	return seededRand().Int63()
}

func DefaultUintFunction() uint {
	return uint(seededRand().Int())
}

func DefaultUint8Function() uint8 {
	return uint8(seededRand().Intn(255))
}

func DefaultUint16Function() uint16 {
	return uint16(seededRand().Intn(65535))
}

func DefaultUint32Function() uint32 {
	return seededRand().Uint32()
}

func DefaultUint64Function() uint64 {
	return seededRand().Uint64()
}

func DefaultFloat32Function() float32 {
	return seededRand().Float32()
}

func DefaultFloat64Function() float64 {
	return seededRand().Float64()
}

func DefaultBoolFunction() bool {
	var val bool
	if seededRand().Intn(1) == 1 {
		val = true
	}
	return val
}

func DefaultStringPointerFunction() *string {
	val := DefaultStringFunction()
	return &val
}

func DefaultIntPointerFunction() *int {
	val := DefaultIntFunction()
	return &val
}

func DefaultInt8PointerFunction() *int8 {
	val := DefaultInt8Function()
	return &val
}

func DefaultInt16PointerFunction() *int16 {
	val := DefaultInt16Function()
	return &val
}

func DefaultInt32PointerFunction() *int32 {
	val := DefaultInt32Function()
	return &val
}

func DefaultInt64PointerFunction() *int64 {
	val := DefaultInt64Function()
	return &val
}

func DefaultUintPointerFunction() *uint {
	val := DefaultUintFunction()
	return &val
}

func DefaultUint8PointerFunction() *uint8 {
	val := DefaultUint8Function()
	return &val
}

func DefaultUint16PointerFunction() *uint16 {
	val := DefaultUint16Function()
	return &val
}

func DefaultUint32PointerFunction() *uint32 {
	val := DefaultUint32Function()
	return &val
}

func DefaultUint64PointerFunction() *uint64 {
	val := DefaultUint64Function()
	return &val
}

func DefaultFloat32PointerFunction() *float32 {
	val := DefaultFloat32Function()
	return &val
}

func DefaultFloat64PointerFunction() *float64 {
	val := DefaultFloat64Function()
	return &val
}

func DefaultBoolPointerFunction() *bool {
	val := DefaultBoolFunction()
	return &val
}
