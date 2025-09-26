package testenv

import (
	"fmt"
	"testing"
)

func TestGoToolPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	path := GoToolPath(t)
	if path == "" {
		t.Error("path is empty")
	}
}
