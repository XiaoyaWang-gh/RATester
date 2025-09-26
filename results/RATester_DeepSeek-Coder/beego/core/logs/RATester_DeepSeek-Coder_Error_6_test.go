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
			name: "TestError",
			args: args{
				f: "TestError",
				v: []interface{}{"TestError"},
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
