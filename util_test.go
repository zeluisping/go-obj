package main

import (
	"testing"
)

func Test_parseFloat(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name     string
		args     args
		wantVal  float32
		wantRest string
		wantOk   bool
	}{
		{"empty value", args{""}, 0, "", false},
		{"not a number", args{"a"}, 0, "", false},
		{"integral", args{"5"}, 5, "", true},
		{"integral + rubish", args{"5foo"}, 5, "foo", true},
		{"start with dot", args{".5"}, 0, "", false},
		{"float", args{"3.1415"}, 3.1415, "", true},
		{"end with dot", args{"5."}, 0, "", false},
		{"double dot", args{"5.3.1"}, 0, "", false},
		{"float range error", args{"400000000000000000000000000000000000000"}, 0, "", false},
		{"negative prefix", args{"-5.423garbage"}, -5.423, "garbage", true},
		{"positive prefix", args{"+5.423garbage"}, +5.423, "garbage", true},
		{"only negative prefix", args{"-"}, 0, "", false},
		{"only positive prefix", args{"+"}, 0, "", false},
		{"only negative prefix garbage", args{"-garbage"}, 0, "", false},
		{"only positive prefix garbage", args{"+garbage"}, 0, "", false},
		{"multiple prefix", args{"-+22"}, 0, "", false},
		{"negative zero", args{"-0"}, 0, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVal, gotRest, gotOk := parseFloat(tt.args.v)
			if gotVal != tt.wantVal {
				t.Errorf("parseFloat() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
			if gotRest != tt.wantRest {
				t.Errorf("parseFloat() gotRest = %v, want %v", gotRest, tt.wantRest)
			}
			if gotOk != tt.wantOk {
				t.Errorf("parseFloat() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_parseInteger(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name     string
		args     args
		wantVal  int32
		wantRest string
		wantOk   bool
	}{
		{"empty value", args{""}, 0, "", false},
		{"not a number", args{"a"}, 0, "", false},
		{"with garbage", args{"15a"}, 15, "a", true},
		{"with dot", args{"6.9bb"}, 6, ".9bb", true},
		{"overflow", args{"4000000000"}, 0, "", false},
		{"negative prefix", args{"-128garbage"}, -128, "garbage", true},
		{"positive prefix", args{"+256"}, 256, "", true},
		{"only prefix", args{"-"}, 0, "", false},
		{"only prefix and garbage", args{"+garbage"}, 0, "", false},
		{"multiple prefix", args{"+-22"}, 0, "", false},
		{"negative zero", args{"-0"}, 0, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVal, gotRest, gotOk := parseInteger(tt.args.v)
			if gotVal != tt.wantVal {
				t.Errorf("parseInteger() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
			if gotRest != tt.wantRest {
				t.Errorf("parseInteger() gotRest = %v, want %v", gotRest, tt.wantRest)
			}
			if gotOk != tt.wantOk {
				t.Errorf("parseInteger() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_skipSpaces(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name     string
		args     args
		wantRest string
		wantOk   bool
	}{
		{"empty value", args{""}, "", false},
		{"missing space", args{"15"}, "", false},
		{"single space", args{" 15"}, "15", true},
		{"single tab", args{"\t15"}, "15", true},
		{"multiple mixed", args{" \t   \t\t 15"}, "15", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRest, gotOk := skipSpaces(tt.args.s)
			if gotRest != tt.wantRest {
				t.Errorf("skipSpaces() gotRest = %v, want %v", gotRest, tt.wantRest)
			}
			if gotOk != tt.wantOk {
				t.Errorf("skipSpaces() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
