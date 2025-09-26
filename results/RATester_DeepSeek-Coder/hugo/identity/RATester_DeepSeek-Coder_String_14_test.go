package identity

import (
	"fmt"
	"testing"
)

func TestString_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	im := &identityManager{
		name: "test",
	}

	expected := "IdentityManager(test)"
	result := im.String()

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
