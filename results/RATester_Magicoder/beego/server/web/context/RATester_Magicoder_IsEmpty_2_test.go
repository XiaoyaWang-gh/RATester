package context

import (
	"fmt"
	"testing"
)

func TestIsEmpty_2(t *testing.T) {
	tests := []struct {
		name   string
		status int
		want   bool
	}{
		{"Status 201", 201, true},
		{"Status 204", 204, true},
		{"Status 304", 304, true},
		{"Status 200", 200, false},
		{"Status 404", 404, false},
		{"Status 500", 500, false},
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
			if got := output.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
