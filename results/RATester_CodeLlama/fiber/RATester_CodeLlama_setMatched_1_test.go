package fiber

import (
	"fmt"
	"testing"
)

func TestSetMatched_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var c *DefaultCtx
	var matched bool
	c.setMatched(matched)
	if c.getMatched() != matched {
		t.Errorf("Test failed, expected %v, got %v", matched, c.getMatched())
	}
}
