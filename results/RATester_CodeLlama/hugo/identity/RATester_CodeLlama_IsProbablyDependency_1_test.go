package identity

import (
	"fmt"
	"testing"
)

func TestIsProbablyDependency_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	id := &predicateIdentity{
		id: "id",
		probablyDependency: func(other Identity) bool {
			return other.IdentifierBase() == "id"
		},
	}

	other := &predicateIdentity{
		id: "other",
	}

	if id.IsProbablyDependency(other) {
		t.Errorf("id.IsProbablyDependency(other) = true, want false")
	}

	other.id = "id"
	if !id.IsProbablyDependency(other) {
		t.Errorf("id.IsProbablyDependency(other) = false, want true")
	}
}
