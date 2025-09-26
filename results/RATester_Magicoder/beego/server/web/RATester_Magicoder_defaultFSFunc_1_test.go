package web

import (
	"fmt"
	"testing"
)

func TestdefaultFSFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := defaultFSFunc()
	if fs == nil {
		t.Error("defaultFSFunc() should not return nil")
	}
}
