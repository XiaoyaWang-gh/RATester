package utils

import (
	"fmt"
	"testing"
)

func TestClear_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	strTo := StrTo("test")
	strTo.Clear()
	if strTo != "" {
		t.Errorf("Expected '', but got %s", strTo)
	}
}
