package web

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/server/web/context/param"
)

func TestGetMethod_1(t *testing.T) {
	type fields struct {
		pattern        string
		controllerType reflect.Type
		methods        map[string]string
		handler        http.Handler
		runFunction    HandleFunc
		routerType     int
		initialize     func() ControllerInterface
		methodParams   []*param.MethodParam
		sessionOn      bool
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
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

			c := &ControllerInfo{
				pattern:        tt.fields.pattern,
				controllerType: tt.fields.controllerType,
				methods:        tt.fields.methods,
				handler:        tt.fields.handler,
				runFunction:    tt.fields.runFunction,
				routerType:     tt.fields.routerType,
				initialize:     tt.fields.initialize,
				methodParams:   tt.fields.methodParams,
				sessionOn:      tt.fields.sessionOn,
			}
			if got := c.GetMethod(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ControllerInfo.GetMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}
