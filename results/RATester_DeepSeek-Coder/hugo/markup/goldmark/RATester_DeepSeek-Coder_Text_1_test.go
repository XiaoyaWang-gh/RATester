package goldmark

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/types/hstring"
)

func TestText_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := headingContext{
		text: hstring.HTML("Test Text"),
	}

	expected := hstring.HTML("Test Text")
	result := ctx.Text()

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
