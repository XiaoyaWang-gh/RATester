package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := Required{Key: "testKey"}
	expected := "testKey"
	actual := r.GetKey()
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
