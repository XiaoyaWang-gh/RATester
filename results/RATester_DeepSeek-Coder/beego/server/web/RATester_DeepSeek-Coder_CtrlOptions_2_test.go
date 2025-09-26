package web

import (
	"fmt"
	"testing"
)

func TestCtrlOptions_2(t *testing.T) {
	type args struct {
		rootpath string
		f        interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Case 1",
			args: args{
				rootpath: "/test",
				f:        "field name",
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

			CtrlOptions(tt.args.rootpath, tt.args.f)
		})
	}
}
