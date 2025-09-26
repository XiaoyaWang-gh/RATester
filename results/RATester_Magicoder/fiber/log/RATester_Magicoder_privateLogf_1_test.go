package log

import (
	"fmt"
	"testing"
)

func TestprivateLogf_1(t *testing.T) {
	type args struct {
		lv      Level
		format  string
		fmtArgs []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			l := &defaultLogger{}
			l.privateLogf(tt.args.lv, tt.args.format, tt.args.fmtArgs)
		})
	}
}
