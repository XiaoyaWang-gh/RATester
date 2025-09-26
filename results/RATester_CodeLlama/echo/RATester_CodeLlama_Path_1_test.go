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

	c := &context{
		path: "/path",
	}
	if got := c.Path(); got != c.path {
		t.Errorf("Path() = %v, want %v", got, c.path)
	}
}
