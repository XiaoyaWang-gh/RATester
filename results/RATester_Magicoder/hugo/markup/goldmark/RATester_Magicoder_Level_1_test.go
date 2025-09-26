package goldmark

import (
	"fmt"
	"testing"
)

func TestLevel_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := headingContext{
		level: 2,
	}

	if ctx.Level() != 2 {
		t.Errorf("Expected level to be 2, but got %d", ctx.Level())
	}
}
