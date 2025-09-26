package web

import (
	"fmt"
	"testing"
)

func TestCtrlPost_4(t *testing.T) {
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
				f:        "TestFunc",
			},
		},
		{
			name: "Test Case 2",
			args: args{
				rootpath: "/test2",
				f:        "TestFunc2",
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

			CtrlPost(tt.args.rootpath, tt.args.f)
		})
	}
}
