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

	if result != "Test Plain Text" {
		t.Errorf("Expected 'Test Plain Text', but got '%s'", result)
	}
}
