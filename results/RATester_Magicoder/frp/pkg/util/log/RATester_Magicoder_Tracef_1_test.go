package log

import (
	"fmt"
	"testing"
)

func TestTracef_1(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Tracef",
			args: args{
				format: "Test Tracef",
				v:      []interface{}{"Test Tracef"},
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

			Tracef(tt.args.format, tt.args.v...)
		})
	}
}
