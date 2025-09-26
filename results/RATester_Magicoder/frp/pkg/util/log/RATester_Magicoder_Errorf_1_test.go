package log

import (
	"fmt"
	"testing"
)

func TestErrorf_1(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Errorf with field name",
			args: args{
				format: "field name: %s",
				v:      []interface{}{"value"},
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

			Errorf(tt.args.format, tt.args.v...)
		})
	}
}
