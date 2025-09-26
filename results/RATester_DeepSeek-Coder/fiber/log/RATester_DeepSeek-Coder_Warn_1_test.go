package log

import (
	"fmt"
	"testing"
)

func TestWarn_1(t *testing.T) {
	tests := []struct {
		name string
		args []any
	}{
		{
			name: "Test case 1",
			args: []any{"Hello", "World"},
		},
		{
			name: "Test case 2",
			args: []any{1, 2, 3},
		},
		{
			name: "Test case 3",
			args: []any{true, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			Warn(tt.args...)
		})
	}
}
