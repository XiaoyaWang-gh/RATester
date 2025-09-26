package testing

import (
	"fmt"
	"testing"
)

func TestgetPort_1(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test Case 1", "8080"},
		{"Test Case 2", "8080"},
		{"Test Case 3", "8080"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getPort(); got != tt.want {
				t.Errorf("getPort() = %v, want %v", got, tt.want)
			}
		})
	}
}
