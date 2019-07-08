package main

import (
	"reflect"
	"testing"
)

func TestVertexNormal_Marshal(t *testing.T) {
	type fields struct {
		I float32
		J float32
		K float32
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
		{"no omit precision -1", fields{7, 12.6236, 0.112156}, args{
			MarshalOptions{
				OmitDefaultOptional: false,
				FloatPrecision:      -1,
			},
		}, "vn 7 12.6236 0.112156"},
		{"no omit precision 6", fields{7, 12.6236, 0.112156}, args{
			MarshalOptions{
				OmitDefaultOptional: false,
				FloatPrecision:      6,
			},
		}, "vn 7.000000 12.623600 0.112156"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vn := &VertexNormal{
				I: tt.fields.I,
				J: tt.fields.J,
				K: tt.fields.K,
			}
			if got := vn.Marshal(tt.args.options); got != tt.want {
				t.Errorf("VertexNormal.Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnmarshalVertexNormal(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantVn  VertexNormal
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVn, err := UnmarshalVertexNormal(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalVertexNormal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotVn, tt.wantVn) {
				t.Errorf("UnmarshalVertexNormal() = %v, want %v", gotVn, tt.wantVn)
			}
		})
	}
}
