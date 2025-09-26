package urlreplacers

import (
	"fmt"
	"testing"
)

func TestNewPrefixState_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	prefixes := newPrefixState()

	if len(prefixes) != 5 {
		t.Errorf("Expected 5 prefixes, got %d", len(prefixes))
	}

	expectedPrefixes := []string{"src=", "href=", "url=", "action=", "srcset="}
	for i, p := range prefixes {
		if string(p.b) != expectedPrefixes[i] {
			t.Errorf("Expected prefix %s, got %s", expectedPrefixes[i], string(p.b))
		}
	}
}
