package source

import (
	"fmt"
	"testing"
)

func TestIsZero_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	g := GitInfo{}
	if !g.IsZero() {
		t.Errorf("expected %v to be zero", g)
	}
}
