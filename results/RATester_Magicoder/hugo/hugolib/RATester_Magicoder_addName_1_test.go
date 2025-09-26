package hugolib

import (
	"fmt"
	"testing"
)

func TestaddName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	handler := &shortcodeHandler{
		nameSet: make(map[string]bool),
	}

	handler.addName("testName")

	if _, ok := handler.nameSet["testName"]; !ok {
		t.Errorf("Expected name to be in the set, but it was not.")
	}
}
