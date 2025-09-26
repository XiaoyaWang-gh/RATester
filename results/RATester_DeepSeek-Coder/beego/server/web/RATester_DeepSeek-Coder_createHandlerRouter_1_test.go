package web

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCreateHandlerRouter_1(t *testing.T) {
	type args struct {
		h       http.Handler
		pattern string
	}
	tests := []struct {
		name string
		args args
		want *ControllerInfo
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

			p := &ControllerRegister{}
			if got := p.createHandlerRouter(tt.args.h, tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ControllerRegister.createHandlerRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
