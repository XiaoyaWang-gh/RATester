package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestAddRouter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{}
	tree.AddRouter("/test", "test")

	ctx := &context.Context{}
	runObject := tree.Match("/test", ctx)

	if runObject != "test" {
		t.Errorf("Expected 'test', got '%v'", runObject)
	}
}
