package crypt

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPKCS5Padding_1(t *testing.T) {
	type args struct {
		ciphertext []byte
		blockSize  int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Test Case 1",
			args: args{
				ciphertext: []byte("hello"),
				blockSize:  8,
			},
			want: []byte("hello\x03\x03\x03"),
		},
		{
			name: "Test Case 2",
			args: args{
				ciphertext: []byte("world"),
				blockSize:  16,
			},
			want: []byte("world\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := PKCS5Padding(tt.args.ciphertext, tt.args.blockSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PKCS5Padding() = %v, want %v", got, tt.want)
			}
		})
	}
}
