package helpers

import (
	"fmt"
	"testing"
)

func TestGetBaseURLRoot_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &PathSpec{}
	path := ""
	if got := p.getBaseURLRoot(path); got != "" {
		t.Errorf("getBaseURLRoot() = %v, want %v", got, "")
	}
}
