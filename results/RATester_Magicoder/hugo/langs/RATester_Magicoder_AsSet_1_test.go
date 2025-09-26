package langs

import (
	"fmt"
	"testing"
)

func TestAsSet_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	languages := Languages{
		{Lang: "en"},
		{Lang: "no"},
		{Lang: "en"},
	}

	set := languages.AsSet()

	if len(set) != 2 {
		t.Errorf("Expected set to have 2 elements, got %d", len(set))
	}

	if !set["en"] {
		t.Error("Expected set to contain 'en'")
	}

	if !set["no"] {
		t.Error("Expected set to contain 'no'")
	}
}
