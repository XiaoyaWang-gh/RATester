package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context/param"
)

func TestAddWithMethodParams_1(t *testing.T) {
	type args struct {
		pattern      string
		c            ControllerInterface
		methodParams []*param.MethodParam
		opts         []ControllerOption
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
			p.addWithMethodParams(tt.args.pattern, tt.args.c, tt.args.methodParams, tt.args.opts...)
		})
	}
}
