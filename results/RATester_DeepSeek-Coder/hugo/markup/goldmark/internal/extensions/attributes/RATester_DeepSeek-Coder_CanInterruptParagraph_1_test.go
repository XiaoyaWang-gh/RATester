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
	expected := true
	actual := parser.CanInterruptParagraph()
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
