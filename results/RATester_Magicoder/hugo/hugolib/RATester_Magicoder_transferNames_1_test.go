package hugolib

import (
	"fmt"
	"reflect"
	"testing"
)

func TesttransferNames_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	handler1 := &shortcodeHandler{
		nameSet: map[string]bool{
			"name1": true,
			"name2": true,
		},
	}
	handler2 := &shortcodeHandler{
		nameSet: map[string]bool{
			"name3": true,
			"name4": true,
		},
	}
	handler1.transferNames(handler2)
	expected := map[string]bool{
		"name1": true,
		"name2": true,
		"name3": true,
		"name4": true,
	}
	if !reflect.DeepEqual(handler1.nameSet, expected) {
		t.Errorf("Expected %v, but got %v", expected, handler1.nameSet)
	}
}
