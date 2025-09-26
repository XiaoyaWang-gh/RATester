package logs

import (
	"fmt"
	"testing"
)

func TestInfo_2(t *testing.T) {
	type args struct {
		f interface{}
		v []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestInfo",
			args: args{
				f: "TestInfo",
				v: []interface{}{"TestInfo"},
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

			Info(tt.args.f, tt.args.v...)
		})
	}
}
