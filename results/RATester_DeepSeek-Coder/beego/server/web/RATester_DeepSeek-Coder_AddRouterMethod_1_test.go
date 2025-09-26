package web

import (
	"fmt"
	"testing"
)

func TestAddRouterMethod_1(t *testing.T) {
	type args struct {
		httpMethod string
		pattern    string
		f          interface{}
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

			p := &ControllerRegister{}
			p.AddRouterMethod(tt.args.httpMethod, tt.args.pattern, tt.args.f)
		})
	}
}
