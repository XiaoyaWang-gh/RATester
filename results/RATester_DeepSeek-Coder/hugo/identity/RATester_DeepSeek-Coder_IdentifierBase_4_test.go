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

	id := &predicateIdentity{
		id: "testID",
	}

	expected := "testID"
	result := id.IdentifierBase()

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
