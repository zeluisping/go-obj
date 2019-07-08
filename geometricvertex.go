package main

import (
	"strconv"
	"strings"
)

// GeometricVertex _ (v)
type GeometricVertex struct {
	X, Y, Z float32
	W       float32
}

// Marshal _
func (v *GeometricVertex) Marshal(options MarshalOptions) string {
	var sb strings.Builder

	sb.WriteString("v ")
	sb.WriteString(strconv.FormatFloat(float64(v.X), 'f', options.FloatPrecision, 32))
	sb.WriteRune(' ')
	sb.WriteString(strconv.FormatFloat(float64(v.Y), 'f', options.FloatPrecision, 32))
	sb.WriteRune(' ')
	sb.WriteString(strconv.FormatFloat(float64(v.Z), 'f', options.FloatPrecision, 32))

	if options.OmitDefaultOptional && v.W == 1 {
		return sb.String()
	}

	sb.WriteRune(' ')
	sb.WriteString(strconv.FormatFloat(float64(v.W), 'f', options.FloatPrecision, 32))

	return sb.String()
}

// ErrVertexBadFieldX _
// var ErrVertexBadFieldX = errors.New("bad field x")
// // ErrVertexBadFieldY _
// var ErrVertexBadFieldY = errors.New("bad field y")
// // ErrVertexBadFieldZ _
// var ErrVertexBadFieldZ = errors.New("bad field y")

// UnmarshalVertex _
func UnmarshalVertex(v string) (vertex GeometricVertex, rest string, ok bool) {
	if v == "" {
		ok = false
		return
	}
	if v[0] != 'v' {
		ok = false
		return
	}
	v, ok = skipSpaces(v[1:])
	if !ok {
		return
	}
	x, v, ok := parseFloat(v)
	if !ok {
		return
	}
	v, ok = skipSpaces(v)
	if !ok {
		return
	}
	y, v, ok := parseFloat(v)
	if !ok {
		return
	}
	v, ok = skipSpaces(v)
	if !ok {
		return
	}
	z, v, ok := parseFloat(v)
	if !ok {
		return
	}
	if v == "" {
		vertex = GeometricVertex{x, y, z, 0}
		return
	}
	v, ok = skipSpaces(v)
	if !ok {
		return
	}
	if v == "" {
		vertex = GeometricVertex{x, y, z, 0}
		return
	}
	w, v, ok := parseFloat(v)
	if !ok {
		return
	}
	rest = v
	vertex = GeometricVertex{x, y, z, w}
	return
}
