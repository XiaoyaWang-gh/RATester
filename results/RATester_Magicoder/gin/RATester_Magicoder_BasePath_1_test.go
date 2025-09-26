package gin

import (
	"fmt"
	"testing"
)

func TestBasePath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	group := &RouterGroup{
		basePath: "/test",
	}

	expected := "/test"
	actual := group.BasePath()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
