package web

import (
	"fmt"
	"testing"
)

func TestCtrlPatch_3(t *testing.T) {
	type args struct {
		rootpath string
		f        interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			CtrlPatch(tt.args.rootpath, tt.args.f)
		})
	}
}
