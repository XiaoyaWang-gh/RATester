package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_19(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	zip := ZipCode{
		Key: "12345",
	}

	expected := "12345"
	actual := zip.GetKey()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
