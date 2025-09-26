package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNSDelete_1(t *testing.T) {
	type args struct {
		rootpath string
		f        HandleFunc
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

			if got := NSDelete(tt.args.rootpath, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NSDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}
