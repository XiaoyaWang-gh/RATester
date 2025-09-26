package goldmark

import (
	"fmt"
	"testing"
)

func TestDestination_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := linkContext{
		destination: "https://example.com",
	}

	expected := "https://example.com"
	actual := ctx.Destination()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
