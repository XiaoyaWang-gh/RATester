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
		{"301", 301, true},
		{"302", 302, true},
		{"303", 303, true},
		{"307", 307, true},
		{"200", 200, false},
		{"404", 404, false},
		{"500", 500, false},
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
				t.Errorf("BeegoOutput.IsRedirect() = %v, want %v", got, tt.want)
			}
		})
	}
}
