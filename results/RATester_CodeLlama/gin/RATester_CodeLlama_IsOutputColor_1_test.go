package gin

import (
	"fmt"
	"testing"
)

func TestIsOutputColor_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &LogFormatterParams{
		isTerm: true,
	}
	if !p.IsOutputColor() {
		t.Errorf("IsOutputColor() = %v, want %v", false, true)
	}
}
