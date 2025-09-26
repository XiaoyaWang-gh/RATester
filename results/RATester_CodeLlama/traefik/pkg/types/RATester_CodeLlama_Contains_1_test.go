package types

import (
	"fmt"
	"testing"
)

func TestContains_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := HTTPCodeRanges{
		{200, 299},
		{400, 499},
		{500, 599},
	}

	if !h.Contains(200) {
		t.Errorf("expected true, got false")
	}

	if h.Contains(199) {
		t.Errorf("expected false, got true")
	}

	if h.Contains(600) {
		t.Errorf("expected false, got true")
	}
}
