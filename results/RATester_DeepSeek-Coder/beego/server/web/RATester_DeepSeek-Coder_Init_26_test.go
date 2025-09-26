package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestInit_26(t *testing.T) {
	type args struct {
		ctx            *context.Context
		controllerName string
		actionName     string
		app            interface{}
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

			c := &Controller{}
			c.Init(tt.args.ctx, tt.args.controllerName, tt.args.actionName, tt.args.app)
		})
	}
}
