package logs

import (
	"fmt"
	"testing"
)

func TestCritical_2(t *testing.T) {
	type args struct {
		f interface{}
		v []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestCritical",
			args: args{
				f: "TestCritical",
				v: []interface{}{"TestCritical"},
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

			Critical(tt.args.f, tt.args.v...)
		})
	}
}
