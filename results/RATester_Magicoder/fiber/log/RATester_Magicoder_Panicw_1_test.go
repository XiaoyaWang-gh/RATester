package log

import (
	"fmt"
	"testing"
)

func TestPanicw_1(t *testing.T) {
	type args struct {
		msg           string
		keysAndValues []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				msg:           "Test panic",
				keysAndValues: []any{"key1", "value1", "key2", "value2"},
			},
		},
		{
			name: "Test case 2",
			args: args{
				msg:           "Test panic with no values",
				keysAndValues: []any{},
			},
		},
		{
			name: "Test case 3",
			args: args{
				msg:           "Test panic with nil values",
				keysAndValues: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			Panicw(tt.args.msg, tt.args.keysAndValues...)
		})
	}
}
