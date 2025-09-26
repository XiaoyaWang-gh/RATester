package validation

import (
	"fmt"
	"testing"
)

func TestGetKey_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := Email{
		Key: "test",
	}
	if e.GetKey() != "test" {
		t.Errorf("Expected %s, got %s", "test", e.GetKey())
	}
}
