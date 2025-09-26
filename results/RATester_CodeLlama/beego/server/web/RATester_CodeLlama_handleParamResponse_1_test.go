package web

import (
	"fmt"
	"reflect"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestHandleParamResponse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var p *ControllerRegister
	var context *beecontext.Context
	var execController ControllerInterface
	var results []reflect.Value
	p.handleParamResponse(context, execController, results)
}
