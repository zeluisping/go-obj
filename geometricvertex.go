package main

import (
	"regexp"
	"strconv"
	"strings"
)

var regGeometricVertexFormat *regexp.Regexp

func init() {
	regGeometricVertexFormat = regexp.MustCompile(`^([-+]?[0-9]*(?:\.?[0-9]+)) (?:([-+]?[0-9]*(?:\.?[0-9]+)) ([-+]?[0-9]*(?:\.?[0-9]+))?)?$`)
}

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
	// if v.W != 1 {
	sb.WriteRune(' ')
	sb.WriteString(strconv.FormatFloat(float64(v.W), 'f', options.FloatPrecision, 32))
	// }

	return sb.String()
}

// UnmarshalVertex _
func UnmarshalVertex(s string) (*GeometricVertex, error) {
	match := regGeometricVertexFormat.FindStringSubmatch(s)
	if match == nil {
		return nil, ErrInvalidFormat
	}

	match = match[1:]

	// ensure right number of arguments
	args, _, err := parseFloatArguments(match, 3, 4)
	if err != nil {
		return nil, err
	}

	for len(args) < 4 {
		args = append(args, 1)
	}

	return &GeometricVertex{args[0], args[1], args[2], args[3]}, nil
}
