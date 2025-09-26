package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPost_4(t *testing.T) {
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

			if got := Post(tt.args.rootpath, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Post() = %v, want %v", got, tt.want)
			}
		})
	}
}
