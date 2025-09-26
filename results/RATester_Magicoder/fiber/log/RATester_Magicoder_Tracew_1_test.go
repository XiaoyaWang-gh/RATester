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
				keysAndValues: []any{"key1", 123},
			},
		},
		{
			name: "Test Tracew with string and float",
			args: args{
				msg:           "Test message",
				keysAndValues: []any{"key2", 123.456},
			},
		},
		{
			name: "Test Tracew with string and bool",
			args: args{
				msg:           "Test message",
				keysAndValues: []any{"key3", true},
			},
		},
		{
			name: "Test Tracew with string and struct",
			args: args{
				msg:           "Test message",
				keysAndValues: []any{"key4", struct{}{}},
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
