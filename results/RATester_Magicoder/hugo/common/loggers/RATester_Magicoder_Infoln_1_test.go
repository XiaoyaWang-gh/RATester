package loggers

import (
	"fmt"
	"testing"
)

func TestInfoln_1(t *testing.T) {
	tests := []struct {
		name string
		args []any
	}{
		{
			name: "Test case 1",
			args: []any{"test1", 123, true},
		},
		{
			name: "Test case 2",
			args: []any{"test2", 456, false},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			l := &logAdapter{}
			l.Infoln(tt.args...)
			// Add assertions here to verify the expected behavior
		})
	}
}
