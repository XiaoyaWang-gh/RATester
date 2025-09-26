package admin

import (
	"fmt"
	"testing"
)

func TestGet_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := make(moduleCommands)
	m["test"] = &doNothingCommand{}
	c := m.Get("test")
	if c == nil {
		t.Errorf("Get failed")
	}
	c = m.Get("test2")
	if c != nil {
		t.Errorf("Get failed")
	}
}
