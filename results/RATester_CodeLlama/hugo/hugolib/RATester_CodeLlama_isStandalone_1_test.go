package hugolib

import (
	"fmt"
	"testing"
)

func TestIsStandalone_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageMeta{}
	if p.isStandalone() {
		t.Error("isStandalone() = true, want false")
	}
}
