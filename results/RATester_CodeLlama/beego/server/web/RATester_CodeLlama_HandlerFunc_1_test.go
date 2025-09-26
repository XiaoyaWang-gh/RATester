package web

import (
	"fmt"
	"testing"
)

func TestHandlerFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Controller{}
	c.methodMapping = make(map[string]func())
	c.methodMapping["fnname"] = func() {}
	if !c.HandlerFunc("fnname") {
		t.Error("HandlerFunc failed")
	}
}
