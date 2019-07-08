package main

import (
	"strconv"
	"strings"
)

// TextureVertex _ (vt)
type TextureVertex struct {
	U float32
	V float32
	W float32
}

// Marshal _
func (tc *TextureVertex) Marshal(options MarshalOptions) string {
	var sb strings.Builder

	sb.WriteString("vt ")
	sb.WriteString(strconv.FormatFloat(float64(tc.U), 'f', options.FloatPrecision, 32))
	if options.OmitDefaultOptional && tc.V == 0 && tc.W == 0 {
		return sb.String()
	}
	sb.WriteRune(' ')
	sb.WriteString(strconv.FormatFloat(float64(tc.V), 'f', options.FloatPrecision, 32))
	if options.OmitDefaultOptional && tc.W == 0 {
		return sb.String()
	}
	sb.WriteRune(' ')
	sb.WriteString(strconv.FormatFloat(float64(tc.W), 'f', options.FloatPrecision, 32))

	return sb.String()
}

// UnmarshalTextureVertex _
func UnmarshalTextureVertex(v string) (vt TextureVertex, rest string, ok bool) {
	if v == "" {
		ok = false
		return
	}
	if len(v) < 2 || v[0] != 'v' || v[1] != 't' {
		ok = false
		return
	}
	v, ok = skipSpaces(v[2:])
	if !ok {
		return
	}
	x, v, ok := parseFloat(v)
	if !ok {
		return
	}
	if v == "" {
		vt = TextureVertex{x, 0, 0}
		return
	}
	v, ok = skipSpaces(v)
	if !ok {
		return
	}
	if v == "" {
		vt = TextureVertex{x, 0, 0}
		return
	}
	y, v, ok := parseFloat(v)
	if !ok {
		return
	}
	if v == "" {
		vt = TextureVertex{x, y, 0}
		return
	}
	v, ok = skipSpaces(v)
	if !ok {
		return
	}
	if v == "" {
		vt = TextureVertex{x, y, 0}
		return
	}
	z, v, ok := parseFloat(v)
	if !ok {
		return
	}
	rest = v
	vt = TextureVertex{x, y, z}
	return
}

// func UnmarshalTextureVertex(s string) (*TextureVertex, error) {
// 	match := regTextureVertexFormat.FindStringSubmatch(s)
// 	if match == nil {
// 		return nil, ErrInvalidFormat
// 	}

// 	match = match[1:]

// 	args, _, err := parseFloatArguments(match, 1, 3)
// 	if err != nil {
// 		return nil, ErrInvalidArgument
// 	}

// 	for len(args) < 3 {
// 		args = append(args, 0)
// 	}

// 	return &TextureVertex{args[0], args[1], args[2]}, nil
// }
