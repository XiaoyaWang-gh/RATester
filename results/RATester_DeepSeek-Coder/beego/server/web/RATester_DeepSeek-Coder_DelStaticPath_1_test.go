package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDelStaticPath_1(t *testing.T) {
	type args struct {
		url string
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

			if got := DelStaticPath(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DelStaticPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
