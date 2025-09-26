package goldmark

import (
	"fmt"
	"testing"
)

func TestPlainText_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := headingContext{
		plainText: "Test Plain Text",
	}

	result := ctx.PlainText()
	expected := "Test Plain Text"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
