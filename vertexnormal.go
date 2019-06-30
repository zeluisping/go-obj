package main

import (
	"regexp"
	"strconv"
	"strings"
)

var regVertexNormalFormat *regexp.Regexp

func init() {
	regVertexNormalFormat = regexp.MustCompile(`^([-+]?[0-9]*(?:\.?[0-9]+)) ([-+]?[0-9]*(?:\.?[0-9]+)) ([-+]?[0-9]*(?:\.?[0-9]+))$`)
}

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
func UnmarshalVertexNormal(s string) (*VertexNormal, error) {
	match := regVertexNormalFormat.FindStringSubmatch(s)
	if match == nil {
		return nil, ErrInvalidFormat
	}

	match = match[1:]

	args, _, err := parseFloatArguments(match, 3, 3)
	if err != nil {
		return nil, ErrInvalidArgument
	}

	return &VertexNormal{args[0], args[1], args[2]}, nil
}
