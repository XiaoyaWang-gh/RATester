package log

import (
	"fmt"
	"testing"
)

func TestInfow_1(t *testing.T) {
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
				msg:           "Test message",
				keysAndValues: []any{"key1", "value1", "key2", "value2"},
			},
		},
		{
			name: "Test case 2",
			args: args{
				msg:           "Another test message",
				keysAndValues: []any{"key3", "value3", "key4", "value4"},
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

			Infow(tt.args.msg, tt.args.keysAndValues...)
		})
	}
}
