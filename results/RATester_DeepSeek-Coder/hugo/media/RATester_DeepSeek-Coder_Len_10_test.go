package media

import (
	"fmt"
	"testing"
)

func TestLen_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	types := Types{
		{Type: "application/json"},
		{Type: "application/xml"},
		{Type: "text/html"},
	}

	expected := 3
	actual := types.Len()

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
