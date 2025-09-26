package identity

import (
	"fmt"
	"testing"
)

func TestIdentifierBase_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	id := &predicateIdentity{id: "test"}
	if id.IdentifierBase() != "test" {
		t.Errorf("IdentifierBase() = %v, want %v", id.IdentifierBase(), "test")
	}
}
