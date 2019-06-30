package main

// Object structure with three components
type Object struct {
	Name string

	Vertices  []*GeometricVertex
	TexCoords []*TextureVertex
	Normals   []*VertexNormal

	Faces   []*Face
	Indices []*Index
}

// Marshal _
func (f *Object) Marshal(options MarshalOptions) (string, error) {
	return "", nil
}

// UnmarshalObject _
func UnmarshalObject(lines []string) (*Object, error) {
	return nil, nil
}
