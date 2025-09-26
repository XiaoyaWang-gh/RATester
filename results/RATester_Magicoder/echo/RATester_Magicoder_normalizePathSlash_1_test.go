package echo

import (
	"fmt"
	"testing"
)

func TestNormalizePathSlash_1(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{"empty string", "", "/"},
		{"starts with slash", "/test", "/test"},
		{"does not start with slash", "test", "/test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := normalizePathSlash(tt.path); got != tt.want {
				t.Errorf("normalizePathSlash() = %v, want %v", got, tt.want)
			}
		})
	}
}
