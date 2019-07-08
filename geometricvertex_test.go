package main

import (
	"reflect"
	"testing"
)

func TestGeometricVertex_Marshal(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
		W float32
	}
	type args struct {
		options MarshalOptions
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"no OmitDefaultOptional", fields{0, 0, 0, 1}, args{
			MarshalOptions{
				OmitDefaultOptional: false,
				FloatPrecision:      6,
			},
		}, "v 0.000000 0.000000 0.000000 1.000000"},
		{"with OmitDefaultOptional", fields{0, 0, 0, 1}, args{
			MarshalOptions{
				OmitDefaultOptional: true,
				FloatPrecision:      6,
			},
		}, "v 0.000000 0.000000 0.000000"},
		// @todo way more cases, will basically test the compliance with the options
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &GeometricVertex{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
				W: tt.fields.W,
			}
			if got := v.Marshal(tt.args.options); got != tt.want {
				t.Errorf("GeometricVertex.Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnmarshalVertex(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name       string
		args       args
		wantVertex GeometricVertex
		wantRest   string
		wantOk     bool
	}{
		{"empty value", args{""}, GeometricVertex{}, "", false},
		{"missing prefix", args{"0 0 0"}, GeometricVertex{}, "", false},
		{"missing x prefix", args{"v0 0 0"}, GeometricVertex{}, "", false},
		{"x,y,z", args{"v 9 7 3"}, GeometricVertex{9, 7, 3, 0}, "", true},
		{"x,y,z trailing spaces", args{"v 9 7 3 \t "}, GeometricVertex{9, 7, 3, 0}, "", true},
		{"x,y,z weird spacing", args{"v 6.12\t\t6.7     1 \t\t\t"}, GeometricVertex{6.12, 6.7, 1, 0}, "", true},
		{"x y z w", args{"v 1 -2 3 4"}, GeometricVertex{1, -2, 3, 4}, "", true},
		{"x y z w trailing", args{"v 1 2 3 4  "}, GeometricVertex{1, 2, 3, 4}, "  ", true},
		{"x y z w trailing garbage", args{"v 1 2 3 4    s    "}, GeometricVertex{1, 2, 3, 4}, "    s    ", true},
		{"x y z wgarbage", args{"v 1 2 3 4garbage"}, GeometricVertex{1, 2, 3, 4}, "garbage", true},
		{"invalid x", args{"v a 2 3 4"}, GeometricVertex{}, "", false},
		{"invalid y", args{"v 1 b 3 4"}, GeometricVertex{}, "", false},
		{"invalid z", args{"v 1 2 c 4"}, GeometricVertex{}, "", false},
		{"invalid w", args{"v 1 2 3 d"}, GeometricVertex{}, "", false},
		{"invalid separator x-y", args{"v 1/2 3 4"}, GeometricVertex{}, "", false},
		{"invalid separator y-z", args{"v 1 2/3 4"}, GeometricVertex{}, "", false},
		{"invalid separator z-w", args{"v 1 2 3/4"}, GeometricVertex{}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVertex, gotRest, gotOk := UnmarshalVertex(tt.args.v)
			if !reflect.DeepEqual(gotVertex, tt.wantVertex) {
				t.Errorf("UnmarshalVertex() gotVertex = %v, want %v", gotVertex, tt.wantVertex)
			}
			if gotRest != tt.wantRest {
				t.Errorf("UnmarshalVertex() gotRest = %v, want %v", gotRest, tt.wantRest)
			}
			if gotOk != tt.wantOk {
				t.Errorf("UnmarshalVertex() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
