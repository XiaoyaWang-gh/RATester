package web

import (
	"fmt"
	"testing"
)

func TestaddPrefix_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{}
	tree.AddRouter("/test", "test")
	addPrefix(tree, "/prefix")

	runObject := tree.Match("/prefix/test", nil)
	if runObject != "test" {
		t.Errorf("Expected 'test', got '%s'", runObject)
	}
}
