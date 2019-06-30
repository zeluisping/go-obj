package main

import (
	"fmt"
	"testing"
)

var marshalTests = []struct {
	Options MarshalOptions
	Vertex  GeometricVertex
	Output  string
}{
	{
		MarshalOptions{ /*FloatPrecision: 0*/ },
		GeometricVertex{321.123456789, -0.987, 0.9876, 1.98765},
		"v 321 -1 1 2",
	},
	// @note these are two cases where the X component will have
	// a weird value, this is a consequence of how floats work.
	// {
	// 	MarshalOptions{FloatPrecision: -1},
	// 	Vertex{321.9876, -0.98765, 0.987654, 1.987654321},
	// 	"v 321.9876 -0.98765 0.987654 1.9876543",
	// },
	// {
	// 	MarshalOptions{FloatPrecision: -1},
	// 	Vertex{321.987654321, -0.98765, 0.987654, 1.987654321},
	// 	"v 321.9876 -0.98765 0.987654 1.9876543",
	// },
	{
		MarshalOptions{FloatPrecision: 3},
		GeometricVertex{321.98765432, -0.987, 0.9876, 1.987654},
		"v 321.988 -0.987 0.988 1.988",
	},
	{
		MarshalOptions{FloatPrecision: 3},
		GeometricVertex{321.987, -0.987, 0.9876, 1.98765},
		"v 321.987 -0.987 0.988 1.988",
	},
	{
		MarshalOptions{FloatPrecision: 6},
		GeometricVertex{1, 0, -0, -1},
		"v 1.000000 0.000000 0.000000 -1.000000",
	},
	{
		MarshalOptions{FloatPrecision: 7},
		GeometricVertex{1, 0, -0, -1},
		"v 1.0000000 0.0000000 0.0000000 -1.0000000",
	},
}

const errFormatTestVertexMarshalFail = `
test-id   := %d
precision := %d 
expected  := "%s"
result    := "%s"`

func TestVertexMarshal(t *testing.T) {
	for i, test := range marshalTests {
		name := fmt.Sprintf("%d-%+v.Marshal()", i, test.Vertex)

		t.Run(name, func(t *testing.T) {
			r := test.Vertex.Marshal(test.Options)

			if r == test.Output {
				return
			}

			t.Errorf(errFormatTestVertexMarshalFail, i, test.Options.FloatPrecision, test.Output, r)
		})
	}
}

var unmarshalTests = []struct {
	Arguments []string
	Expected  *GeometricVertex
	Error     error
}{
	{
		[]string{},
		nil,
		ErrInsuficientArguments,
	},
	{
		[]string{"1", "2", "3", "4", "5"},
		nil,
		ErrTooManyArguments,
	},
	{
		[]string{"1", "2", "3", "a"},
		nil,
		ErrInvalidArgument,
	},
	{
		[]string{"1", "1", "1"},
		&GeometricVertex{1, 1, 1, 1},
		nil,
	},
	{
		[]string{"1", "1", "1", "-7"},
		&GeometricVertex{1, 1, 1, -7},
		nil,
	},
	{
		[]string{"0.123456789", "1000.987654321", "-321.987654321"},
		&GeometricVertex{0.123456789, 1000.987654321, -321.987654321, 1},
		nil,
	},
}

var errFormatTestVertexUnmarshalFailError = `
arguments      	:= %+v
expected-error	:= %v
result-error	:= %v
`

var errFormatTestVertexUnmarshalFailValue = `
arguments		:= %+v
expected-value 	:= %+v
result-value	:= %+v
`

func TestVertexUnmarshal(t *testing.T) {
	for i, test := range unmarshalTests {
		name := fmt.Sprintf("%d-UnmarshalVertex(%+v)", i, test.Arguments)

		t.Run(name, func(t *testing.T) {
			v, err := UnmarshalVertex(test.Arguments)
			if err != test.Error {
				t.Errorf(errFormatTestVertexUnmarshalFailError, test.Arguments, test.Error, err)
				return
			}

			e := test.Expected

			if v == nil && v == e {
				return
			}

			if v == nil || e == nil {
				t.Errorf(errFormatTestVertexUnmarshalFailValue, test.Arguments, test.Expected, v)
				return
			}

			if v.X == e.X && v.Y == e.Y && v.Z == e.Z && v.W == e.W {
				return
			}

			t.Errorf(errFormatTestVertexUnmarshalFailValue, test.Arguments, test.Expected, v)
		})
	}
}
