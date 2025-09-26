package loggers

import (
	"fmt"
	"testing"
)

func TestWarnidf_2(t *testing.T) {
	type args struct {
		id     string
		format string
		v      []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				id:     "test_id",
				format: "test_format",
				v:      []any{"test_value"},
			},
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
			l.Warnidf(tt.args.id, tt.args.format, tt.args.v...)
			// Add assertions to check the expected behavior
		})
	}
}
