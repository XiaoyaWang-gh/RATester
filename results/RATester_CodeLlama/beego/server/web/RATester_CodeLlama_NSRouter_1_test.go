package web

import (
	"fmt"
	"testing"
)

func TestNSRouter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var rootpath string
	var c ControllerInterface
	var mappingMethods []string
	var linkNamespace LinkNamespace

	rootpath = "rootpath"
	c = &Controller{}
	mappingMethods = []string{"GET", "POST"}
	linkNamespace = NSRouter(rootpath, c, mappingMethods...)

	if linkNamespace == nil {
		t.Errorf("NSRouter return nil")
	}
}
