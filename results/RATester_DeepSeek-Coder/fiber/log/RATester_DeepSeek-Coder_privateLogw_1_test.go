package log

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestPrivateLogw_1(t *testing.T) {
	type args struct {
		lv            Level
		format        string
		keysAndValues []any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test privateLogw with LevelInfo, format and keysAndValues",
			args: args{
				lv:            LevelInfo,
				format:        "This is a test",
				keysAndValues: []any{"key1", "value1", "key2", "value2"},
			},
			want: "INFO This is a test key1=value1 key2=value2",
		},
		{
			name: "Test privateLogw with LevelDebug, no format and keysAndValues",
			args: args{
				lv:            LevelDebug,
				format:        "",
				keysAndValues: []any{},
			},
			want: "DEBUG",
		},
		{
			name: "Test privateLogw with LevelError, format and keysAndValues",
			args: args{
				lv:            LevelError,
				format:        "This is another test",
				keysAndValues: []any{"key3", "value3", "key4", "value4"},
			},
			want: "ERROR This is another test key3=value3 key4=value4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			l := &defaultLogger{
				stdlog: log.New(os.Stdout, "", log.LstdFlags),
				level:  tt.args.lv,
			}
			l.privateLogw(tt.args.lv, tt.args.format, tt.args.keysAndValues)
			// TODO: Add assertions to check the output of the logger
		})
	}
}
