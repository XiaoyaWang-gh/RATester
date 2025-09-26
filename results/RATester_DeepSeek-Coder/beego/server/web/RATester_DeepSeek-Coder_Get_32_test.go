package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGet_32(t *testing.T) {
	type args struct {
		rootpath string
		f        HandleFunc
	}
	tests := []struct {
		name string
		n    *Namespace
		args args
		want *Namespace
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

			if got := tt.n.Get(tt.args.rootpath, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Namespace.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
