package web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &Namespace{}
	rootpath := "rootpath"
	h := http.Handler(nil)
	n.Handler(rootpath, h)
}
