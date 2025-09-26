package plugins

import (
	"fmt"
	"testing"
)

func TestGoPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{
		goPath: "/path/to/go",
	}

	expected := "/path/to/go"
	actual := client.GoPath()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
