package web

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestHandler_3(t *testing.T) {
	type args struct {
		rootpath string
		h        http.Handler
		options  []interface{}
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

			if got := Handler(tt.args.rootpath, tt.args.h, tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler() = %v, want %v", got, tt.want)
			}
		})
	}
}
