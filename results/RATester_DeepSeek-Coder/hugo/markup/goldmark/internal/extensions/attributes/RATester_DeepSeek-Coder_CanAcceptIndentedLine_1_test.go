package attributes

import (
	"fmt"
	"testing"
)

func TestCanAcceptIndentedLine_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	parser := &attrParser{}
	result := parser.CanAcceptIndentedLine()
	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}
