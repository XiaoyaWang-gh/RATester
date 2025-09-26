package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNSCtrlPut_1(t *testing.T) {
	type args struct {
		rootpath string
		f        interface{}
	}
	tests := []struct {
		name string
		args args
		want LinkNamespace
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

			if got := NSCtrlPut(tt.args.rootpath, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NSCtrlPut() = %v, want %v", got, tt.want)
			}
		})
	}
}
