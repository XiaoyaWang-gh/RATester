package web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &ControllerRegister{}
	pattern := "pattern"
	h := http.Handler(nil)
	options := []interface{}{}
	p.Handler(pattern, h, options...)
}
