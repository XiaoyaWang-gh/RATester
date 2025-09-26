package log

import (
	"fmt"
	"testing"
)

func TestPrivateLog_1(t *testing.T) {
	type args struct {
		lv      Level
		fmtArgs []any
	}
	tests := []struct {
		name string
		l    *defaultLogger
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

			tt.l.privateLog(tt.args.lv, tt.args.fmtArgs)
		})
	}
}
