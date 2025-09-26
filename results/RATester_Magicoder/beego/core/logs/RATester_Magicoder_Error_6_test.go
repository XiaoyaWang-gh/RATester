package logs

import (
	"fmt"
	"testing"
)

func TestError_6(t *testing.T) {
	type args struct {
		f interface{}
		v []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestError_0",
			args: args{
				f: "test",
				v: []interface{}{"test"},
			},
		},
		{
			name: "TestError_1",
			args: args{
				f: "test",
				v: []interface{}{"test", 1},
			},
		},
		{
			name: "TestError_2",
			args: args{
				f: "test",
				v: []interface{}{"test", 1, 2.0},
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

			Error(tt.args.f, tt.args.v...)
		})
	}
}
