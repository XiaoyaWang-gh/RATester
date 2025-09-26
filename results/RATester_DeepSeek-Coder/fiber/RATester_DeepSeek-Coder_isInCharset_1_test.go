package fiber

import (
	"fmt"
	"testing"
)

func TestIsInCharset_1(t *testing.T) {
	charset := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	tests := []struct {
		name       string
		searchChar byte
		want       bool
	}{
		{"Test1", 'a', true},
		{"Test2", 'A', true},
		{"Test3", '0', true},
		{"Test4", '9', true},
		{"Test5", ' ', false},
		{"Test6", '@', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isInCharset(tt.searchChar, charset); got != tt.want {
				t.Errorf("isInCharset() = %v, want %v", got, tt.want)
			}
		})
	}
}
