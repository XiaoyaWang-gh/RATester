package hugolib

import (
	"fmt"
	"testing"
)

func TestIsPage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageMeta{}
	if !p.IsPage() {
		t.Errorf("IsPage() = false, want true")
	}
}
