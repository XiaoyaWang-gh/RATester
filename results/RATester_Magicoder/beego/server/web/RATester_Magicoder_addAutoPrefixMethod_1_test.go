package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestaddAutoPrefixMethod_1(t *testing.T) {
	type args struct {
		prefix         string
		controllerName string
		methodName     string
		ctrl           reflect.Type
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
			p.addAutoPrefixMethod(tt.args.prefix, tt.args.controllerName, tt.args.methodName, tt.args.ctrl)
		})
	}
}
