package hugolib

import (
	"fmt"
	"testing"
)

func TestHasName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &shortcodeHandler{}
	s.nameSet = make(map[string]bool)
	s.nameSet["name"] = true
	if !s.hasName("name") {
		t.Errorf("hasName failed")
	}
}
