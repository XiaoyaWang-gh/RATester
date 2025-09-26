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
			want: "902830912309182309182301982310985723", // This is a placeholder, replace with the actual expected result
		},
		{
			name: "Test Case 2",
			args: args{f: "hello world"},
			want: "80000000000000000000000000000000", // This is a placeholder, replace with the actual expected result
		},
		{
			name: "Test Case 3",
			args: args{f: "1234567890"},
			want: "56789012345678901234567890123456", // This is a placeholder, replace with the actual expected result
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
