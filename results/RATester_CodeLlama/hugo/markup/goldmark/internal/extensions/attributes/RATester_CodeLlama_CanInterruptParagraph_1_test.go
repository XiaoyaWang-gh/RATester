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

	a := &attrParser{}
	if got := a.CanInterruptParagraph(); got != true {
		t.Errorf("attrParser.CanInterruptParagraph = %v, want %v", got, true)
	}
}
