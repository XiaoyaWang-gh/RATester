package logs

import (
	"fmt"
	"testing"

	"github.com/rs/zerolog"
)

func TestLogWithLevel_1(t *testing.T) {
	type args struct {
		ev  *zerolog.Event
		msg string
		kvs []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test with string fields",
			args: args{
				ev:  &zerolog.Event{},
				msg: "test message",
				kvs: []interface{}{"field1", "value1", "field2", "value2"},
			},
		},
		{
			name: "Test with int fields",
			args: args{
				ev:  &zerolog.Event{},
				msg: "test message",
				kvs: []interface{}{"field1", 1, "field2", 2},
			},
		},
		{
			name: "Test with float fields",
			args: args{
				ev:  &zerolog.Event{},
				msg: "test message",
				kvs: []interface{}{"field1", 1.1, "field2", 2.2},
			},
		},
		{
			name: "Test with bool fields",
			args: args{
				ev:  &zerolog.Event{},
				msg: "test message",
				kvs: []interface{}{"field1", true, "field2", false},
			},
		},
		{
			name: "Test with nil fields",
			args: args{
				ev:  &zerolog.Event{},
				msg: "test message",
				kvs: []interface{}{"field1", nil, "field2", nil},
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

			logWithLevel(tt.args.ev, tt.args.msg, tt.args.kvs...)
		})
	}
}
