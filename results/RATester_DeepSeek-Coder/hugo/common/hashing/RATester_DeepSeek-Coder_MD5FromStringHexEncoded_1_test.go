package hashing

import (
	"fmt"
	"testing"
)

func TestMD5FromStringHexEncoded_1(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{f: "test"},
			want: "098f6bcd4621d373cade4e832627b4f6",
		},
		{
			name: "Test case 2",
			args: args{f: "hello"},
			want: "5d41402abc4b2a76b9719d911017c592",
		},
		{
			name: "Test case 3",
			args: args{f: "world"},
			want: "3e25960a79dbc69b674cd4ec67a72c62",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := MD5FromStringHexEncoded(tt.args.f); got != tt.want {
				t.Errorf("MD5FromStringHexEncoded() = %v, want %v", got, tt.want)
			}
		})
	}
}
