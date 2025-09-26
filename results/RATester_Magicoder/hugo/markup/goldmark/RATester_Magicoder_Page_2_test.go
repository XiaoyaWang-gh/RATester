package goldmark

import (
	"fmt"
	"testing"
)

func TestPage_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := headingContext{
		page: "test_page",
	}

	expected := "test_page"
	result := ctx.Page()

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
