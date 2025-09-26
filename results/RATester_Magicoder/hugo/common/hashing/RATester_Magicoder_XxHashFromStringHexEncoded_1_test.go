package hashing

import (
	"fmt"
	"testing"
)

func TestXxHashFromStringHexEncoded_1(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{f: "test"},
			want: "902fbdd2b1df0c4f",
		},
		{
			name: "Test Case 2",
			args: args{f: "hello"},
			want: "309ecc489c12d6eb",
		},
		{
			name: "Test Case 3",
			args: args{f: "world"},
			want: "4e03657aea95a47d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := XxHashFromStringHexEncoded(tt.args.f); got != tt.want {
				t.Errorf("XxHashFromStringHexEncoded() = %v, want %v", got, tt.want)
			}
		})
	}
}
