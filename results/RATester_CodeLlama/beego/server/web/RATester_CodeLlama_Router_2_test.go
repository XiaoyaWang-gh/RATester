package web

import (
	"fmt"
	"testing"
)

func TestRouter_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &Namespace{}
	rootpath := "rootpath"
	c := &Controller{}
	mappingMethods := []string{"GET", "POST"}
	n.Router(rootpath, c, mappingMethods...)
}
