package factogo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ModelA struct {
	String                    string
	Int                       int
	Int8                      int8
	Int16                     int16
	Int32                     int32
	Int64                     int64
	Uint                      uint
	Uint8                     uint8
	Uint16                    uint16
	Uint32                    uint32
	Uint64                    uint64
	Float32                   float32
	Float64                   float64
	Bool                      bool
	ModelB                    ModelB
	IntArray                  []int
	ModelBArray               []ModelB
	StringMap                 map[string]string
	IntPointerArray           []*int
	ModelBPointerArray        []*ModelB
	StringPointerMap          map[string]*string
	StringPointer             *string
	IntPointer                *int
	Int8Pointer               *int8
	Int16Pointer              *int16
	Int32Pointer              *int32
	Int64Pointer              *int64
	UintPointer               *uint
	Uint8Pointer              *uint8
	Uint16Pointer             *uint16
	Uint32Pointer             *uint32
	Uint64Pointer             *uint64
	Float32Pointer            *float32
	Float64Pointer            *float64
	BoolPointer               *bool
	ModelC                    *ModelC
	IntArrayPointer           *[]int
	ModelCArrayPointer        *[]ModelC
	StringMapPointer          *map[string]string
	IntPointerArrayPointer    *[]*int
	ModelCPointerArrayPointer *[]*ModelC
	StringPointerMapPointer   *map[string]*string
}

type ModelB struct {
	String string
}

type ModelC struct {
	String string
}

type ModelD struct {
	String string
	Int    int
}

func TestDesign(t *testing.T) {
	Clear()
	assert.NoError(t, Factory("admin").Design(&ModelA{}), "Should be designed")
	assert.NotNil(t, factories["admin"], "Factory design should be added")
}

func TestDesignUsingDefaultFunctions(t *testing.T) {
	Clear()
	assert.NoError(t, Factory("admin").
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
		Set("Bool").
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
		Set("BoolPointer").
		Set("IntArrayPointer").
		Set("ModelCArrayPointer").
		Set("StringMapPointer").
		Set("IntPointerArrayPointer").
		Set("ModelCPointerArrayPointer").
		Set("StringPointerMapPointer").
		Design(&ModelA{}),
		"Should be designed")
	assert.NotNil(t, factories["admin"], "Factory design should be added")
}

func TestDesignUsingDefaultFunctionsWithComplexTypes(t *testing.T) {
	Clear()
	assert.NoError(t, Factory("admin").
		Set("ModelB").
		Design(&ModelA{}),
		"Should be designed")
	assert.NotNil(t, factories["admin"], "Factory design should be added")
}

func TestDesignUsingSet(t *testing.T) {
	Clear()
	assert.NoError(t,
		Factory("admin").Set("String", "My string").Design(&ModelA{}),
		"Should be designed")
	assert.NotNil(t, factories["admin"], "Factory design should be added")
}

func TestDesignUsingSetWithFunction(t *testing.T) {
	Clear()
	function := func() string { return "My string" }
	assert.NoError(t,
		Factory("admin").Set("String", function).Design(&ModelA{}),
		"Should be designed")
	assert.NotNil(t, factories["admin"], "Factory design should be added")
}

func TestDesignUsingSetOfUnexistingFieldName(t *testing.T) {
	Clear()
	assert.EqualErrorf(t, Factory("admin").Set("Unexisting", nil).Design(&ModelA{}),
		"The struct 'ModelA' doesn't have a field 'Unexisting'",
		"Shoud return error")
	assert.Nil(t, factories["admin"], "Factory design should not be added")
}

func TestDesignUsingSetWithWrongFieldType(t *testing.T) {
	Clear()
	assert.EqualErrorf(t, Factory("admin").Set("String", 10).Design(&ModelA{}),
		"The type of value of 'String' not match the type in the struct ModelA",
		"Shoud return error")
	assert.Nil(t, factories["admin"], "Factory design should not be added")
}

func TestDesignUsingSetWithWrongFieldTypeWithFunction(t *testing.T) {
	Clear()

	function := func() int { return 10 }

	assert.EqualErrorf(t, Factory("admin").Set("String", function).Design(&ModelA{}),
		"The type of value of 'String' not match the type in the struct ModelA",
		"Shoud return error")
	assert.Nil(t, factories["admin"], "Factory design should not be added")
}

func TestDesignUsingAnAlreadyDesignedFactory(t *testing.T) {
	Clear()

	Factory("modelA").Design(&ModelA{})

	assert.Errorf(t, Factory("modelA").Design(&ModelA{}),
		"A designed Factory named 'modelA' already exists")
}
