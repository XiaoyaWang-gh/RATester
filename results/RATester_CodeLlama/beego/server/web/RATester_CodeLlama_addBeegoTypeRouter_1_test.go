package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddBeegoTypeRouter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &ControllerRegister{}
	ct := reflect.TypeOf(ControllerRegister{})
	ctMethod := "addBeegoTypeRouter"
	httpMethod := "GET"
	pattern := "/"
	p.addBeegoTypeRouter(ct, ctMethod, httpMethod, pattern)
}
