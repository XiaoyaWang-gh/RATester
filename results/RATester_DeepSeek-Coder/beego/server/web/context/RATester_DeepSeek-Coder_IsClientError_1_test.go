package context

import (
	"fmt"
	"testing"
)

func TestIsClientError_1(t *testing.T) {
	tests := []struct {
		name   string
		status int
		want   bool
	}{
		{"400 status", 400, true},
		{"404 status", 404, true},
		{"499 status", 499, true},
		{"500 status", 500, false},
		{"501 status", 501, false},
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
			if got := output.IsClientError(); got != tt.want {
				t.Errorf("IsClientError() = %v, want %v", got, tt.want)
			}
		})
	}
}
