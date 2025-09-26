package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPut_14(t *testing.T) {
	type args struct {
		rootpath string
		f        HandleFunc
	}
	tests := []struct {
		name string
		args args
		want *HttpServer
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

			if got := Put(tt.args.rootpath, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Put() = %v, want %v", got, tt.want)
			}
		})
	}
}
