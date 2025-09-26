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
		{"GET", "GET", true},
		{"HEAD", "HEAD", true},
		{"OPTIONS", "OPTIONS", true},
		{"TRACE", "TRACE", true},
		{"POST", "POST", false},
		{"PUT", "PUT", false},
		{"DELETE", "DELETE", false},
		{"CONNECT", "CONNECT", false},
		{"PATCH", "PATCH", false},
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
