package log

import (
	"fmt"
	"testing"
)

func TestInitLogger_1(t *testing.T) {
	type args struct {
		logPath         string
		levelStr        string
		maxDays         int
		disableLogColor bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestInitLogger_0",
			args: args{
				logPath:         "console",
				levelStr:        "info",
				maxDays:         7,
				disableLogColor: false,
			},
		},
		{
			name: "TestInitLogger_1",
			args: args{
				logPath:         "test.log",
				levelStr:        "error",
				maxDays:         3,
				disableLogColor: true,
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

			InitLogger(tt.args.logPath, tt.args.levelStr, tt.args.maxDays, tt.args.disableLogColor)
		})
	}
}
