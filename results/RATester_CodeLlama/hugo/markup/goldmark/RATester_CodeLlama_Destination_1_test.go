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

	ctx := linkContext{}
	ctx.destination = "https://www.google.com"
	if ctx.Destination() != "https://www.google.com" {
		t.Errorf("Destination() = %v, want %v", ctx.Destination(), "https://www.google.com")
	}
}
