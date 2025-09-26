package web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNSHandler_1(t *testing.T) {
	t.Run("Test NSHandler", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var rootpath string
		var h http.Handler
		var ns *Namespace
		var linkNamespace LinkNamespace
		linkNamespace = NSHandler(rootpath, h)
		ns = &Namespace{}
		linkNamespace(ns)
	})
}
