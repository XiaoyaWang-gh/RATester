package log

import (
	"fmt"
	"testing"
)

func TestWarnf_1(t *testing.T) {
	tests := []struct {
		name   string
		format string
		v      []any
	}{
		{
			name:   "Test case 1",
			format: "This is a test",
			v:      []any{"test"},
		},
		{
			name:   "Test case 2",
			format: "This is another test",
			v:      []any{"another", "test"},
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

			Warnf(tt.format, tt.v...)
			// Add assertions here to verify the expected behavior
		})
	}
}
