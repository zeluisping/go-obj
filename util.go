package main

import (
	"strconv"
)

func vIndex(vertices []*GeometricVertex, v *GeometricVertex) (int, error) {
	for i, x := range vertices {
		if x == v {
			return i + 1, nil
		}
	}

	return 0, ErrVerticeNotFound
}

func vtIndex(texcoords []*TextureVertex, vt *TextureVertex) (int, error) {
	for i, x := range texcoords {
		if x == vt {
			return i + 1, nil
		}
	}

	return 0, ErrTexCoordNotFound
}

func vnIndex(normals []*VertexNormal, vn *VertexNormal) (int, error) {
	for i, x := range normals {
		if x == vn {
			return i + 1, nil
		}
	}

	return 0, ErrNormalNotFound
}

func assertArgumentCount(length, min, max int) error {
	if length < min {
		return ErrInsuficientArguments
	}
	if max != -1 && length > max {
		return ErrTooManyArguments
	}
	return nil
}

func parseFloatArguments(src []string, min int, max int) ([]float32, int, error) {
	err := assertArgumentCount(len(src), min, max)
	if err != nil {
		return nil, 0, err
	}

	r := make([]float32, 0, len(src))

	// ensure values are valid
	for i, s := range src {
		v, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return nil, i, err
		}

		r = append(r, float32(v))
	}

	return r, len(r) - 1, nil
}
