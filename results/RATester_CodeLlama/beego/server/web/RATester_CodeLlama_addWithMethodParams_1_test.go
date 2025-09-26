package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context/param"
)

func TestAddWithMethodParams_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &ControllerRegister{}
	pattern := "pattern"
	c := &Controller{}
	methodParams := []*param.MethodParam{}
	opts := []ControllerOption{}
	p.addWithMethodParams(pattern, c, methodParams, opts...)
}
