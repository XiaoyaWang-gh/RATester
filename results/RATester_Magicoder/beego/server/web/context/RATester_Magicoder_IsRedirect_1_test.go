package context

import (
	"fmt"
	"testing"
)

func TestIsRedirect_1(t *testing.T) {
	tests := []struct {
		name   string
		status int
		want   bool
	}{
		{"Test 1", 301, true},
		{"Test 2", 302, true},
		{"Test 3", 303, true},
		{"Test 4", 307, true},
		{"Test 5", 200, false},
		{"Test 6", 404, false},
		{"Test 7", 500, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			output := &BeegoOutput{
				Status: tt.status,
			}
			if got := output.IsRedirect(); got != tt.want {
				t.Errorf("IsRedirect() = %v, want %v", got, tt.want)
			}
		})
	}
}
