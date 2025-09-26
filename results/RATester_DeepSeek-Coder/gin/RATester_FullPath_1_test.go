package gin

import (
	"fmt"
	"testing"
)

func TestFullPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	expected := "/test/path"
	context := &Context{
		fullPath: expected,
	}

	result := context.FullPath()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
