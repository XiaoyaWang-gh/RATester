package log

import (
	"fmt"
	"testing"
)

func TestDebugf_1(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestDebugf_0",
			args: args{
				format: "Test Debugf",
				v:      []interface{}{"Test Debugf"},
			},
		},
		{
			name: "TestDebugf_1",
			args: args{
				format: "Test Debugf",
				v:      []interface{}{"Test Debugf"},
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

			Debugf(tt.args.format, tt.args.v...)
		})
	}
}
