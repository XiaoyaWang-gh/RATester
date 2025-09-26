package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := Enum{
		Rules: "test",
		Key:   "testKey",
	}

	key := e.GetKey()
	if key != "testKey" {
		t.Errorf("Expected 'testKey', got '%s'", key)
	}
}
