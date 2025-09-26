package log

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestPrivateLog_1(t *testing.T) {
	type args struct {
		lv      Level
		fmtArgs []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test privateLog with LevelDebug and string message",
			args: args{
				lv:      LevelDebug,
				fmtArgs: []any{"This is a debug message"},
			},
		},
		{
			name: "Test privateLog with LevelInfo and int message",
			args: args{
				lv:      LevelInfo,
				fmtArgs: []any{12345},
			},
		},
		{
			name: "Test privateLog with LevelWarn and float64 message",
			args: args{
				lv:      LevelWarn,
				fmtArgs: []any{123.45},
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

			l := &defaultLogger{
				stdlog: log.New(os.Stdout, "", log.LstdFlags),
				level:  tt.args.lv,
			}
			l.privateLog(tt.args.lv, tt.args.fmtArgs)
		})
	}
}
