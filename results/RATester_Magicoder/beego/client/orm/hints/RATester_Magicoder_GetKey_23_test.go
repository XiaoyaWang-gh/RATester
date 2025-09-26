package hints

import (
	"fmt"
	"testing"
)

func TestGetKey_23(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hint := Hint{
		key:   "testKey",
		value: "testValue",
	}

	key := hint.GetKey()

	if key != "testKey" {
		t.Errorf("Expected 'testKey', but got '%v'", key)
	}
}
