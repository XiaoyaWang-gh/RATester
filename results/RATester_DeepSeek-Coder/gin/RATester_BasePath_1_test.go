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

	expected := "/test"
	group := &RouterGroup{
		basePath: expected,
	}

	result := group.BasePath()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
