package urlreplacers

import (
	"fmt"
	"testing"
)

func TestnewPrefixState_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	prefixes := newPrefixState()

	if len(prefixes) != 5 {
		t.Errorf("Expected 5 prefixes, got %d", len(prefixes))
	}

	for _, p := range prefixes {
		if p.b == nil {
			t.Errorf("Prefix has nil byte slice")
		}
		if p.f == nil {
			t.Errorf("Prefix has nil function")
		}
	}
}
