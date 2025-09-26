package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_16(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	phone := Phone{Key: "1234567890"}
	expected := "1234567890"
	actual := phone.GetKey()
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
