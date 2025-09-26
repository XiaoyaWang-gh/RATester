package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_20(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := Alpha{Key: "testKey"}
	expected := "testKey"
	result := a.GetKey()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
