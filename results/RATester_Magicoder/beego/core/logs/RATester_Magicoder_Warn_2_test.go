package logs

import (
	"fmt"
	"testing"
)

func TestWarn_2(t *testing.T) {
	type args struct {
		f interface{}
		v []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Warn",
			args: args{
				f: "Test Warn",
				v: []interface{}{"Test Warn"},
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

			Warn(tt.args.f, tt.args.v...)
		})
	}
}
