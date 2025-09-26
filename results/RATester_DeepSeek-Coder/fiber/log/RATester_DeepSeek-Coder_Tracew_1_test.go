package log

import (
	"fmt"
	"testing"
)

func TestTracew_1(t *testing.T) {
	type args struct {
		msg           string
		keysAndValues []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Tracew with string and int",
			args: args{
				msg:           "Test message",
				keysAndValues: []any{"key1", "value1", "key2", 2},
			},
		},
		{
			name: "Test Tracew with float",
			args: args{
				msg:           "Test message",
				keysAndValues: []any{"key1", 1.23},
			},
		},
		{
			name: "Test Tracew with bool",
			args: args{
				msg:           "Test message",
				keysAndValues: []any{"key1", true},
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

			Tracew(tt.args.msg, tt.args.keysAndValues...)
		})
	}
}
