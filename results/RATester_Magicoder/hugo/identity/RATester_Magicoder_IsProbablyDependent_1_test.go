package identity

import (
	"fmt"
	"testing"
)

func TestIsProbablyDependent_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	id := &predicateIdentity{
		id: "test",
		probablyDependent: func(other Identity) bool {
			return other.IdentifierBase() == "test"
		},
	}

	other := &predicateIdentity{
		id: "test",
	}

	if !id.IsProbablyDependent(other) {
		t.Error("Expected IsProbablyDependent to return true")
	}
}
