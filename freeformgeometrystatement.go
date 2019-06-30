package main

import (
	"regexp"
	"strconv"
	"strings"
)

var regFreeFormGeometryStatementFormat *regexp.Regexp

func init() {
	regFreeFormGeometryStatementFormat = regexp.MustCompile(`^([-+]?[0-9]*(?:\.?[0-9]+)) (?:([-+]?[0-9]*(?:\.?[0-9]+)) ([-+]?[0-9]*(?:\.?[0-9]+))?)?$`)
}

// FreeFormGeometryStatement _
type FreeFormGeometryStatement struct {
	U float32
	V float32
	W float32
}

// Marshal _
func (vp *FreeFormGeometryStatement) Marshal(options *MarshalOptions) (string, error) {
	var sb strings.Builder

	sb.WriteString("vp ")
	sb.WriteString(strconv.FormatFloat(float64(vp.U), 'f', options.FloatPrecision, 32))
	sb.WriteRune(' ')
	sb.WriteString(strconv.FormatFloat(float64(vp.V), 'f', options.FloatPrecision, 32))
	sb.WriteRune(' ')
	sb.WriteString(strconv.FormatFloat(float64(vp.W), 'f', options.FloatPrecision, 32))

	return sb.String(), nil
}

// UnmarshalFreeFormGeometryStatement _
func UnmarshalFreeFormGeometryStatement(s string) (*FreeFormGeometryStatement, error) {
	match := regFreeFormGeometryStatementFormat.FindStringSubmatch(s)
	if match == nil {
		return nil, ErrInvalidFormat
	}

	args, _, err := parseFloatArguments(match[1:], 1, 3)
	if err != nil {
		// n tells us which value was wrong
		return nil, err
	}

	for len(args) < 3 {
		args = append(args, 1)
	}

	return &FreeFormGeometryStatement{args[0], args[1], args[2]}, nil
}
