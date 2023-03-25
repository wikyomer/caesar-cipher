package main

import (
	"reflect"
	"testing"
)

func Test_caesar(t *testing.T) {
	type args struct {
		text []byte
		key  int8
	}
	tests := []struct {
		name           string
		args           args
		wantCiphertext []byte
	}{
		{"test1", args{[]byte("abc"), 1}, []byte("bcd")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCiphertext := caesar(tt.args.text, tt.args.key); !reflect.DeepEqual(gotCiphertext, tt.wantCiphertext) {
				t.Errorf("caesar() = %v, want %v", gotCiphertext, tt.wantCiphertext)
			}
		})
	}
}

func Test_decoder(t *testing.T) {
	type args struct {
		ciphertext []byte
		key        int8
	}
	tests := []struct {
		name     string
		args     args
		wantText []byte
	}{
		{"test2", args{[]byte("bcdefghijklmnopqrstuvwxyza"), 1}, []byte("abcdefghijklmnopqrstuvwxyz")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotText := decoder(tt.args.ciphertext, tt.args.key); !reflect.DeepEqual(gotText, tt.wantText) {
				t.Errorf("decoder() = %v, want %v", gotText, tt.wantText)
			}
		})
	}
}

func Test_shift(t *testing.T) {
	type args struct {
		originalSlice []byte
		key           int8
	}
	tests := []struct {
		name              string
		args              args
		wantModifiedSlice []byte
	}{
		{"test3", args{[]byte("bcdefghijklmnopqrstuvwxyza"), -1}, []byte("abcdefghijklmnopqrstuvwxyz")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotModifiedSlice := shift(tt.args.originalSlice, tt.args.key); !reflect.DeepEqual(gotModifiedSlice, tt.wantModifiedSlice) {
				t.Errorf("shift() = %v, want %v", gotModifiedSlice, tt.wantModifiedSlice)
			}
		})
	}
}
