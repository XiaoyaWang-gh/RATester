package paths

import (
	"fmt"
	"testing"
)

func TestIdentifierAsString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{
		s: "a/b/c",
	}

	if got := p.identifierAsString(0); got != "" {
		t.Errorf("identifierAsString() = %v, want %v", got, "")
	}

	if got := p.identifierAsString(1); got != "" {
		t.Errorf("identifierAsString() = %v, want %v", got, "")
	}

	if got := p.identifierAsString(2); got != "" {
		t.Errorf("identifierAsString() = %v, want %v", got, "")
	}
}
