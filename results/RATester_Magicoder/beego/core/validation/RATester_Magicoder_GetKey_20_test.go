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
	result := a.GetKey()
	if result != "testKey" {
		t.Errorf("Expected 'testKey', but got '%s'", result)
	}
}
