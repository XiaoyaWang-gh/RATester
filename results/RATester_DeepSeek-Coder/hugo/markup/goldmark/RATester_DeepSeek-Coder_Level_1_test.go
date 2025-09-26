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
		level: 3,
	}

	expected := 3
	actual := ctx.Level()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
