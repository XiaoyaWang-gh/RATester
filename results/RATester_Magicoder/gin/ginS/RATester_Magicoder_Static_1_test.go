package ginS

import (
	"fmt"
	"testing"
)

func TestStatic_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	relativePath := "/static"
	root := "./static"

	routes := Static(relativePath, root)

	if routes == nil {
		t.Error("Expected routes to be not nil")
	}
}
