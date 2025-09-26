package gin

import (
	"fmt"
	"testing"
)

func TestIterate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	path := "/"
	method := "GET"
	routes := RoutesInfo{}
	root := &node{}
	routes = iterate(path, method, routes, root)
	if len(routes) != 0 {
		t.Errorf("Testiterate failed")
	}
}
