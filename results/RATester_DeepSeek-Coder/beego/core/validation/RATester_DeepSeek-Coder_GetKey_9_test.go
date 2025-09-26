package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Mobile{
		Match: Match{},
		Key:   "testKey",
	}

	expected := "testKey"
	actual := m.GetKey()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
