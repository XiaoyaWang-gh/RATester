package goldmark

import (
	"fmt"
	"testing"
)

func TestOrdinal_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := imageLinkContext{
		linkContext: linkContext{},
		ordinal:     1,
		isBlock:     true,
	}

	expected := 1
	actual := ctx.Ordinal()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
