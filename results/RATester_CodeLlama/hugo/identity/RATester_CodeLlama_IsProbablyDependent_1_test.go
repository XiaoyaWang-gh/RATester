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
		id: "id",
		probablyDependent: func(other Identity) bool {
			return other.IdentifierBase() == "id"
		},
	}

	other := &predicateIdentity{
		id: "other",
	}

	if id.IsProbablyDependent(other) {
		t.Errorf("id.IsProbablyDependent(other) = true, want false")
	}

	other.probablyDependent = func(other Identity) bool {
		return other.IdentifierBase() == "id"
	}

	if !id.IsProbablyDependent(other) {
		t.Errorf("id.IsProbablyDependent(other) = false, want true")
	}
}
