package template

import (
	"fmt"
	"testing"
)

func TestisJSIdentPart_1(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{"Test dollar sign", '$', true},
		{"Test number", '5', true},
		{"Test uppercase letter", 'A', true},
		{"Test underscore", '_', true},
		{"Test lowercase letter", 'a', true},
		{"Test special character", '*', false},
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
