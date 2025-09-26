package template

import (
	"fmt"
	"testing"
)

func TestIsJSIdentPart_1(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{"Test dollar sign", '$', true},
		{"Test number 0", '0', true},
		{"Test number 9", '9', true},
		{"Test uppercase A", 'A', true},
		{"Test uppercase Z", 'Z', true},
		{"Test underscore", '_', true},
		{"Test lowercase a", 'a', true},
		{"Test lowercase z", 'z', true},
		{"Test lowercase g", 'g', true},
		{"Test uppercase G", 'G', false},
		{"Test number 5", '5', true},
		{"Test number 1", '1', false},
		{"Test special character", '@', false},
		{"Test space", ' ', false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isJSIdentPart(tt.r); got != tt.want {
				t.Errorf("isJSIdentPart() = %v, want %v", got, tt.want)
			}
		})
	}
}
