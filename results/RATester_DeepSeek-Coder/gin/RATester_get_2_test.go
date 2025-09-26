package gin

import (
	"fmt"
	"testing"
)

func TestGet_2(t *testing.T) {
	mt := methodTrees{
		{
			method: "GET",
			root: &node{
				path:     "/",
				indices:  "",
				children: []*node{},
			},
		},
		{
			method: "POST",
			root: &node{
				path:     "/",
				indices:  "",
				children: []*node{},
			},
		},
	}

	t.Run("TestGetExistingMethod", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		n := mt.get("GET")
		if n == nil {
			t.Errorf("Expected method 'GET' to exist")
		}
		if n.path != "/" {
			t.Errorf("Expected path '/', got %s", n.path)
		}
	})

	t.Run("TestGetNonExistingMethod", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		n := mt.get("PUT")
		if n != nil {
			t.Errorf("Expected method 'PUT' to not exist, got %v", n)
		}
	})
}
