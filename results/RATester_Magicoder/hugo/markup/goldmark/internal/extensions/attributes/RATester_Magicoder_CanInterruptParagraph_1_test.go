package attributes

import (
	"fmt"
	"testing"
)

func TestCanInterruptParagraph_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	parser := &attrParser{}
	result := parser.CanInterruptParagraph()
	if result != true {
		t.Errorf("Expected true, but got %v", result)
	}
}
