package echo

import (
	"fmt"
	"testing"
)

func TestAddMethod_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &node{}
	method := "GET"
	h := &routeMethod{}
	n.addMethod(method, h)
}
