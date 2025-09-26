package log

import (
	"fmt"
	"testing"
)

func TestInfof_1(t *testing.T) {
	type args struct {
		format string
		v      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestInfof_0",
			args: args{
				format: "test %s",
				v:      []interface{}{"value"},
			},
		},
		{
			name: "TestInfof_1",
			args: args{
				format: "test %d",
				v:      []interface{}{123},
			},
		},
		{
			name: "TestInfof_2",
			args: args{
				format: "test %f",
				v:      []interface{}{123.456},
			},
		},
		{
			name: "TestInfof_3",
			args: args{
				format: "test %t",
				v:      []interface{}{true},
			},
		},
		{
			name: "TestInfof_4",
			args: args{
				format: "test %v",
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

			Infof(tt.args.format, tt.args.v...)
		})
	}
}
