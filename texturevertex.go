package main

import (
	"regexp"
	"strconv"
	"strings"
)

var regTextureVertexFormat *regexp.Regexp

func init() {
	regTextureVertexFormat = regexp.MustCompile(`^([-+]?[0-9]*(?:\.?[0-9]+)) (?:([-+]?[0-9]*(?:\.?[0-9]+)) ([-+]?[0-9]*(?:\.?[0-9]+))?)?$`)
}

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
	sb.WriteRune(' ')
	sb.WriteString(strconv.FormatFloat(float64(tc.V), 'f', options.FloatPrecision, 32))
	sb.WriteRune(' ')
	sb.WriteString(strconv.FormatFloat(float64(tc.W), 'f', options.FloatPrecision, 32))

	return sb.String()
}

// UnmarshalTextureVertex _
func UnmarshalTextureVertex(s string) (*TextureVertex, error) {
	match := regTextureVertexFormat.FindStringSubmatch(s)
	if match == nil {
		return nil, ErrInvalidFormat
	}

	match = match[1:]

	args, _, err := parseFloatArguments(match, 1, 3)
	if err != nil {
		return nil, ErrInvalidArgument
	}

	for len(args) < 3 {
		args = append(args, 0)
	}

	return &TextureVertex{args[0], args[1], args[2]}, nil
}
