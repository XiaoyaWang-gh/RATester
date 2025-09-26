package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetStaticPath_1(t *testing.T) {
	type args struct {
		url  string
		path string
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

			if got := SetStaticPath(tt.args.url, tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetStaticPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
