package fiber

import (
	"fmt"
	"testing"
)

func TestIsMethodSafe_1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{"Test GET", MethodGet, true},
		{"Test HEAD", "HEAD", true},
		{"Test OPTIONS", "OPTIONS", true},
		{"Test TRACE", "TRACE", true},
		{"Test POST", "POST", false},
		{"Test PUT", "PUT", false},
		{"Test DELETE", "DELETE", false},
		{"Test CONNECT", "CONNECT", false},
		{"Test PATCH", "PATCH", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := IsMethodSafe(tt.args); got != tt.want {
				t.Errorf("IsMethodSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}
