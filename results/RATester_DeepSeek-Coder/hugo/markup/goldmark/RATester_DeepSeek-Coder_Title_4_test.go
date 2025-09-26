package goldmark

import (
	"fmt"
	"testing"
)

func TestTitle_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := linkContext{
		title: "Test Title",
	}

	expected := "Test Title"
	actual := ctx.Title()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
