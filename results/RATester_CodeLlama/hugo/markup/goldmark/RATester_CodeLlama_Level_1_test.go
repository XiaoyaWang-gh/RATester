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

	ctx := headingContext{}
	if ctx.Level() != 0 {
		t.Errorf("Level() = %v, want %v", ctx.Level(), 0)
	}
}
