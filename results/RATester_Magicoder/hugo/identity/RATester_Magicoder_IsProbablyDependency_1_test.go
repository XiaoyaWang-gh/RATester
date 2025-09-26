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
		id: "test",
		probablyDependency: func(other Identity) bool {
			return other.IdentifierBase() == "test"
		},
	}

	other := &predicateIdentity{
		id: "test",
	}

	if !id.IsProbablyDependency(other) {
		t.Error("Expected IsProbablyDependency to return true")
	}
}
