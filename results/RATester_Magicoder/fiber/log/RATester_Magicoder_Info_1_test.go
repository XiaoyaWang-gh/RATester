package log

import (
	"fmt"
	"testing"
)

func TestInfo_1(t *testing.T) {
	tests := []struct {
		name string
		args []any
	}{
		{
			name: "Test case 1",
			args: []any{"test", 1, true},
		},
		{
			name: "Test case 2",
			args: []any{"test2", 2, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			Info(tt.args...)
		})
	}
}
