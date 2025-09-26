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
		ordinal: 1,
	}

	if ctx.Ordinal() != 1 {
		t.Errorf("Expected 1, got %d", ctx.Ordinal())
	}
}
