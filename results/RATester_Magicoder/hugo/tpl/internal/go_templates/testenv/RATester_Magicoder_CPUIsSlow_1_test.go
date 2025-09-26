package testenv

import (
	"fmt"
	"testing"
)

func TestCPUIsSlow_1(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"arm", true},
		{"mips", true},
		{"mipsle", true},
		{"mips64", true},
		{"mips64le", true},
		{"wasm", true},
		{"x86", false},
		{"x86_64", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := CPUIsSlow(); got != tt.want {
				t.Errorf("CPUIsSlow() = %v, want %v", got, tt.want)
			}
		})
	}
}
