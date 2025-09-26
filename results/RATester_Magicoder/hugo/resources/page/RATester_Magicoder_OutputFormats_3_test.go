package page

import (
	"fmt"
	"testing"
)

func TestOutputFormats_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	outputFormats := nop.OutputFormats()
	if outputFormats != nil {
		t.Errorf("Expected nil, got %v", outputFormats)
	}
}
