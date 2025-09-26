package echo

import (
	"fmt"
	"testing"
)

func TestPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &context{
		path: "/test",
	}

	expected := "/test"
	actual := ctx.Path()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
