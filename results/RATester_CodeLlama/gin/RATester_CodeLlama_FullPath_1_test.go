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

	c := &Context{
		fullPath: "test",
	}
	if c.FullPath() != "test" {
		t.Errorf("FullPath() = %v, want %v", c.FullPath(), "test")
	}
}
