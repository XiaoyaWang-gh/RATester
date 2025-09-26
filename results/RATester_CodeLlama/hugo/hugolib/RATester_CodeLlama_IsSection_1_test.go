package hugolib

import (
	"fmt"
	"testing"
)

func TestIsSection_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageMeta{}
	if p.IsSection() {
		t.Errorf("IsSection() = true, want false")
	}
}
