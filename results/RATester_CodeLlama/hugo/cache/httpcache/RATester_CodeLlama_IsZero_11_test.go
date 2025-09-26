package httpcache

import (
	"fmt"
	"testing"
)

func TestIsZero_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := PollConfigCompiled{}
	if !p.IsZero() {
		t.Errorf("IsZero() = %v, want %v", p.IsZero(), true)
	}
}
