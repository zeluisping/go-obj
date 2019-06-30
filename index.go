package main

import (
	"strconv"
	"strings"
)

// Index structure with three components
type Index struct {
	Vertex   *GeometricVertex
	TexCoord *TextureVertex
	Normal   *VertexNormal
}

// Marshal the `Index` into the Wavefront .obj syntax.
//
// The resulting `string` is in one of the formats:
//
// `v`
//
// `v/vt`
//
// `v/vt/vn`
//
// `v//vn`
func (i *Index) Marshal(vertices []*GeometricVertex, texcoords []*TextureVertex, normals []*VertexNormal, options MarshalOptions) (string, error) {
	var sb strings.Builder

	if i.Vertex == nil {
		return "", ErrIndexVertexRequired
	}

	if texcoords != nil && i.TexCoord == nil {
		return "", ErrMissingTexCoord
	}

	if normals != nil && i.Normal == nil {
		return "", ErrMissingNormal
	}

	if texcoords == nil && normals == nil {
		vi, err := vIndex(vertices, i.Vertex)
		if err != nil {
			return "", err
		}

		return strconv.FormatInt(int64(vi), 64), nil
	}

	if texcoords != nil && normals == nil {
		vi, err := vIndex(vertices, i.Vertex)
		if err != nil {
			return "", err
		}

		vti, err := vtIndex(texcoords, i.TexCoord)
		if err != nil {
			return "", err
		}

		sb.WriteString(strconv.FormatInt(int64(vi), 64))
		sb.WriteRune('/')
		sb.WriteString(strconv.FormatInt(int64(vti), 64))

		return sb.String(), nil
	}

	if texcoords == nil && normals != nil {
		vi, err := vIndex(vertices, i.Vertex)
		if err != nil {
			return "", err
		}

		vni, err := vnIndex(normals, i.Normal)
		if err != nil {
			return "", err
		}

		sb.WriteString(strconv.FormatInt(int64(vi), 64))
		sb.WriteString("//")
		sb.WriteString(strconv.FormatInt(int64(vni), 64))

		return sb.String(), nil
	}

	// texcoords != nil && normals != nil
	vi, err := vIndex(vertices, i.Vertex)
	if err != nil {
		return "", err
	}

	vti, err := vtIndex(texcoords, i.TexCoord)
	if err != nil {
		return "", err
	}

	vni, err := vnIndex(normals, i.Normal)
	if err != nil {
		return "", err
	}

	sb.WriteString(strconv.FormatInt(int64(vi), 64))
	sb.WriteRune('/')
	sb.WriteString(strconv.FormatInt(int64(vti), 64))
	sb.WriteRune('/')
	sb.WriteString(strconv.FormatInt(int64(vni), 64))

	return sb.String(), nil
}

// UnmarshalIndex from a `string` of the indices. Signature must be either of:
//
// `v`
//
// `v/vt`
//
// `v//vn`
//
// `v/vt/vn`
//
// The indices in the `string` must be parseable into a `int64` through `strconv.ParseInt`
func UnmarshalIndex(vertices []*GeometricVertex, texcoords []*TextureVertex, normals []*VertexNormal, value string) (*Index, error) {
	split := strings.Split(value, "/")
	length := len(split)

	vi, err := strconv.ParseInt(split[0], 10, 64)
	if err != nil || vi < 1 {
		return nil, ErrInvalidArgument
	}

	if int(vi) > len(vertices) {
		return nil, ErrVerticeNotFound
	}

	v := vertices[vi-1]
	var vt *TextureVertex
	var vn *VertexNormal

	if length == 2 || length == 3 && split[1] != "" {
		vti, err := strconv.ParseInt(split[1], 10, 64)
		if err != nil {
			return nil, ErrInvalidArgument
		}

		if int(vti) > len(texcoords) {
			return nil, ErrTexCoordNotFound
		}

		vt = texcoords[vti-1]
	}

	if length == 3 {
		vni, err := strconv.ParseInt(split[1], 10, 64)
		if err != nil || vni < 1 {
			return nil, ErrInvalidArgument
		}

		if int(vni) > len(normals) {
			return nil, ErrNormalNotFound
		}

		vn = normals[vni-1]
	}

	return &Index{
		Vertex:   v,
		TexCoord: vt,
		Normal:   vn,
	}, nil
}
