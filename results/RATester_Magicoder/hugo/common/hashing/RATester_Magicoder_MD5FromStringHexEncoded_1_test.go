package hashing

import (
	"fmt"
	"testing"
)

func TestMD5FromStringHexEncoded_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "Test case 1",
			arg:  "test",
			want: "098f6bcd4621d373cade4e832627b4f6",
		},
		{
			name: "Test case 2",
			arg:  "hello",
			want: "5d41402abc4b2a76b9719d911017c592",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := MD5FromStringHexEncoded(tt.arg); got != tt.want {
				t.Errorf("MD5FromStringHexEncoded() = %v, want %v", got, tt.want)
			}
		})
	}
}
