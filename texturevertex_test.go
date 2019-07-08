package main

import (
	"reflect"
	"testing"
)

func TestTextureVertex_Marshal(t *testing.T) {
	type fields struct {
		U float32
		V float32
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
		{"no OmitDefaultOptional", fields{0, 0, 0}, args{
			MarshalOptions{
				OmitDefaultOptional: false,
				FloatPrecision:      0,
			},
		}, "vt 0 0 0"},
		{"omitDefaltOptional V and W default", fields{0, 0, 0}, args{
			MarshalOptions{
				OmitDefaultOptional: true,
				FloatPrecision:      0,
			},
		}, "vt 0"},
		{"omitDefaltOptional W default", fields{0, 1, 0}, args{
			MarshalOptions{
				OmitDefaultOptional: true,
				FloatPrecision:      0,
			},
		}, "vt 0 1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tc := &TextureVertex{
				U: tt.fields.U,
				V: tt.fields.V,
				W: tt.fields.W,
			}
			if got := tc.Marshal(tt.args.options); got != tt.want {
				t.Errorf("TextureVertex.Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnmarshalTextureVertex(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name     string
		args     args
		wantVt   TextureVertex
		wantRest string
		wantOk   bool
	}{
		{"empty value", args{""}, TextureVertex{}, "", false},
		{"missing prefix", args{"0 0 0"}, TextureVertex{}, "", false},
		{"missing x prefix", args{"vt0 0 0"}, TextureVertex{}, "", false},
		{"u", args{"vt 9"}, TextureVertex{9, 0, 0}, "", true},
		{"ugarbage", args{"vt 9garbage"}, TextureVertex{}, "", false},
		{"u trailing spaces", args{"vt 9\t \t   \t"}, TextureVertex{9, 0, 0}, "", true},
		{"u v", args{"vt 9 7"}, TextureVertex{9, 7, 0}, "", true},
		{"u vgarbage", args{"vt 9 7garbage"}, TextureVertex{}, "", false},
		{"u v trailing spaces", args{"vt 9 7 \t "}, TextureVertex{9, 7, 0}, "", true},
		{"u v weird spacing", args{"vt 6.12\t\t6.7 \t\t\t"}, TextureVertex{6.12, 6.7, 0}, "", true},
		{"u v w", args{"vt 6 7 4"}, TextureVertex{6, 7, 4}, "", true},
		{"u v wgarbage", args{"vt 1 2 3garbage"}, TextureVertex{1, 2, 3}, "garbage", true},
		{"u v w trailing spaces", args{"vt 1 2 3    \t    "}, TextureVertex{1, 2, 3}, "    \t    ", true},
		{"u v w trailing garbage", args{"vt 1 2 3    s    "}, TextureVertex{1, 2, 3}, "    s    ", true},
		{"invalid u", args{"vt a 2 3"}, TextureVertex{}, "", false},
		{"invalid v", args{"vt 1 b 3"}, TextureVertex{}, "", false},
		{"invalid w", args{"vt 1 2 c"}, TextureVertex{}, "", false},
		{"invalid separator u-v", args{"vt 1/2 3"}, TextureVertex{}, "", false},
		{"invalid separator v-w", args{"vt 1 2/3"}, TextureVertex{}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVt, gotRest, gotOk := UnmarshalTextureVertex(tt.args.v)
			if !reflect.DeepEqual(gotVt, tt.wantVt) {
				t.Errorf("UnmarshalTextureVertex() gotVt = %v, want %v", gotVt, tt.wantVt)
			}
			if gotRest != tt.wantRest {
				t.Errorf("UnmarshalTextureVertex() gotRest = %v, want %v", gotRest, tt.wantRest)
			}
			if gotOk != tt.wantOk {
				t.Errorf("UnmarshalTextureVertex() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
