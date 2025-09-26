package web

import (
	"fmt"
	"testing"
)

func TestDefaultFSFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := defaultFSFunc()
	if fs == nil {
		t.Errorf("defaultFSFunc() = %v, want not nil", fs)
	}
}
