package main

import (
	"slices"
	"testing"

	"github.com/gohryt/asphyxia/env"
)

type String struct {
	String string `name:"STRING"`
}

func TestParseString(t *testing.T) {
	env, err := env.Parse[String]()
	if err != nil {
		t.Errorf("%s,%s", "TestParseString: error while Parse()", err.Error())
	}

	if env.String != "http://example.com" {
		t.Errorf("TestParseString: Result was incorrect, got: %s, want: %s.", env.String, "http://example.com")
	}
}

type Int struct {
	Int8  int8  `name:"INT_8"`
	Int16 int16 `name:"INT_16"`
	Int32 int32 `name:"INT_32"`
	Int64 int64 `name:"INT_64"`
}

func TestParseInt(t *testing.T) {
	env, err := env.Parse[Int]()
	if err != nil {
		t.Errorf("%s,%s", "TestParseInt: error while Parse()", err.Error())
	}

	if env.Int8 != int8(127) {
		t.Errorf("TestParseInt: Result was incorrect, got: %d, want: %d.", env.Int8, int8(127))
	}

	if env.Int16 != int16(32767) {
		t.Errorf("TestParseInt: Result was incorrect, got: %d, want: %d.", env.Int16, int16(32767))
	}

	if env.Int32 != int32(2147483647) {
		t.Errorf("TestParseInt: Result was incorrect, got: %d, want: %d.", env.Int32, int32(2147483647))
	}

	if env.Int64 != int64(9223372036854775807) {
		t.Errorf("TestParseInt: Result was incorrect, got: %d, want: %d.", env.Int64, int64(9223372036854775807))
	}
}

type Uint struct {
	Uint8  uint8  `name:"UINT_8"`
	Uint16 uint16 `name:"UINT_16"`
	Uint32 uint32 `name:"UINT_32"`
	Uint64 uint64 `name:"UINT_64"`
}

func TestParseUint(t *testing.T) {
	env, err := env.Parse[Uint]()
	if err != nil {
		t.Errorf("%s,%s", "TestParseUint: error while Parse()", err.Error())
	}

	if env.Uint8 != uint8(255) {
		t.Errorf("TestParseUint: Result was incorrect, got: %d, want: %d.", env.Uint8, uint8(255))
	}

	if env.Uint16 != uint16(65535) {
		t.Errorf("TestParseUint: Result was incorrect, got: %d, want: %d.", env.Uint16, uint16(65535))
	}

	if env.Uint32 != uint32(4294967295) {
		t.Errorf("TestParseUint: Result was incorrect, got: %d, want: %d.", env.Uint32, uint32(4294967295))
	}

	if env.Uint64 != uint64(18446744073709551615) {
		t.Errorf("TestParseUint: Result was incorrect, got: %d, want: %d.", env.Uint64, uint64(18446744073709551615))
	}
}

type Float struct {
	Float32 float32 `name:"FLOAT_32"`
	Float64 float64 `name:"FLOAT_64"`
}

func TestParseFloat(t *testing.T) {
	env, err := env.Parse[Float]()
	if err != nil {
		t.Errorf("%s,%s", "TestParseFloat: error while Parse()", err.Error())
	}

	if env.Float32 != float32(3.14) {
		t.Errorf("TestParseFloat: Result was incorrect, got: %f, want: %f.", env.Float32, float32(3.14))
	}

	if env.Float64 != float64(3.14) {
		t.Errorf("TestParseFloat: Result was incorrect, got: %f, want: %f.", env.Float64, float64(3.14))
	}
}

type Complex struct {
	Complex64  complex64  `name:"COMPLEX_64"`
	Complex128 complex128 `name:"COMPLEX_128"`
}

func TestParseComplex(t *testing.T) {
	env, err := env.Parse[Complex]()
	if err != nil {
		t.Errorf("%s,%s", "TestParseComplex: error while Parse()", err.Error())
	}

	if env.Complex64 != 1+2i {
		t.Errorf("TestParseComplex: Result was incorrect, got: %f, want: %f.", env.Complex64, 1+2i)
	}

	if env.Complex64 != 1+2i {
		t.Errorf("TestParseComplex: Result was incorrect, got: %f, want: %f.", env.Complex128, 1+2i)
	}
}

type Bool struct {
	Bool bool `name:"BOOL"`
}

func TestParseBool(t *testing.T) {
	env, err := env.Parse[Bool]()
	if err != nil {
		t.Errorf("%s,%s", "TestParseBool: error while Parse()", err.Error())
	}

	if !env.Bool {
		t.Errorf("TestParseBool: Result was incorrect, got: %t, want: %t.", env.Bool, true)
	}
}

type Slice struct {
	Slice []string `name:"SLICE"`
}

func TestParseSlice(t *testing.T) {
	env, err := env.Parse[Slice]()
	if err != nil {
		t.Errorf("%s,%s", "TestParseSlice: error while Parse()", err.Error())
	}

	if !slices.Equal(env.Slice, []string{"1", "2", "3", "4", "5"}) {
		t.Errorf("TestParseSlice: Result was incorrect, got: %v, want: %v.", env.Slice, []string{"1", "2", "3", "4", "5"})
	}
}

type ChiledWithDefault struct {
	String struct {
		String string `name:"str" default:"default"`
		Int    struct {
			Int  int `name:"int" default:"-1337"`
			Uint struct {
				Uint uint `name:"uint" default:"420"`
			}
		}
	}
}

func TestChildWithDefaultValue(t *testing.T) {
	env, err := env.Parse[ChiledWithDefault]()
	if err != nil {
		t.Errorf("%s,%s", "TestChildWithDefaultValue: error while Parse()", err.Error())
	}

	if env.String.String != "default" {
		t.Errorf("TestChildWithDefaultValue: Result was incorrect, got: %s, want: %s.", env.String.String, "default")
	}

	if env.String.Int.Int != -1337 {
		t.Errorf("TestChildWithDefaultValue: Result was incorrect, got: %d, want: %d.", env.String.Int.Int, -1337)
	}

	if env.String.Int.Uint.Uint != 420 {
		t.Errorf("TestChildWithDefaultValue: Result was incorrect, got: %d, want: %d.", env.String.Int.Uint.Uint, 420)
	}
}
