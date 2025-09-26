package logs

import (
	"fmt"
	"testing"
)

func TestWarning_2(t *testing.T) {
	type args struct {
		f interface{}
		v []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Warning",
			args: args{
				f: "Test Warning",
				v: []interface{}{"Test Warning"},
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

			Warning(tt.args.f, tt.args.v...)
		})
	}
}
