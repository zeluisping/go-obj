package main

import (
	"strconv"
	"strings"
)

// VertexNormal _ (vn)
type VertexNormal struct {
	I, J, K float32
}

// Marshal _
func (vn *VertexNormal) Marshal(options MarshalOptions) string {
	var sb strings.Builder

	sb.WriteString("vn ")
	sb.WriteString(strconv.FormatFloat(float64(vn.I), 'f', options.FloatPrecision, 32))
	sb.WriteRune(' ')
	sb.WriteString(strconv.FormatFloat(float64(vn.J), 'f', options.FloatPrecision, 32))
	sb.WriteRune(' ')
	sb.WriteString(strconv.FormatFloat(float64(vn.K), 'f', options.FloatPrecision, 32))

	return sb.String()
}

// UnmarshalVertexNormal _
func UnmarshalVertexNormal(s string) (vn VertexNormal, err error) {
	return
}
