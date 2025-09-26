package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_21(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	length := Length{
		N:   5,
		Key: "testKey",
	}

	expected := "testKey"
	actual := length.GetKey()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
