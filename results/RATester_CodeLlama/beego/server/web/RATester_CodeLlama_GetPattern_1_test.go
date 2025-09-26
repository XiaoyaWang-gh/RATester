package web

import (
	"fmt"
	"testing"
)

func TestGetPattern_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ControllerInfo{}
	c.pattern = "test"
	if c.GetPattern() != "test" {
		t.Errorf("GetPattern() = %v, want %v", c.GetPattern(), "test")
	}
}
