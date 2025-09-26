package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := Base64{
		Key: "testKey",
	}

	expected := "testKey"
	actual := b.GetKey()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
