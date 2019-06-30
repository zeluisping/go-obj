package main

import (
	"strings"
)

// Face structure with three components
type Face struct {
	Indices []*Index
}

// Marshal the `Face` into the Wavefront .obj syntax.
//
// The resulting string is in one of the formats:
//
// Without `texcoords` and without `normals`:
//
// `f v1 v2 v3 ...`
//
// With `texcoords`and without `normals`:
//
// `f v1/vt1 v2/vt2 v3/vt3 ...`
//
// Without `texcoords` and with `normals`:
//
// `f v1//vn1 v2//vn2 v3//vn3 ...`
//
// With both `texcoords` and `normals`:
//
// `f v1/vt1/vn1 v2/vt2/vn2 v3/vt3/vn3 ...`
func (f *Face) Marshal(vertices []*GeometricVertex, texcoords []*TextureVertex, normals []*VertexNormal, options MarshalOptions) (string, error) {
	var sb strings.Builder

	sb.WriteRune('f')

	for _, i := range f.Indices {
		s, err := i.Marshal(vertices, texcoords, normals, options)
		if err != nil {
			return "", err
		}

		sb.WriteRune(' ')
		sb.WriteString(s)
	}

	return sb.String(), nil
}

// UnmarshalFace from a slice of `string` arguments.
//
// There must be at least three `Index` per `Face` in order to form a face.
//
// The indices within the arguments must be parseable into a `int64` through `strconv.ParseInt`.
func UnmarshalFace(vertices []*GeometricVertex, texcoords []*TextureVertex, normals []*VertexNormal, arguments []string) (*Face, error) {
	// ensure right number of arguments
	err := assertArgumentCount(len(arguments), 3, -1)
	if err != nil {
		return nil, err
	}

	var f Face

	for _, arg := range arguments {
		i, err := UnmarshalIndex(vertices, texcoords, normals, arg)
		if err != nil {
			return nil, err
		}

		f.Indices = append(f.Indices, i)
	}

	return &f, nil
}
