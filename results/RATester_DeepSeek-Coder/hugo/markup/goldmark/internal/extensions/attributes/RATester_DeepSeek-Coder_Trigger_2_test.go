package attributes

import (
	"bytes"
	"fmt"
	"testing"
)

func TestTrigger_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	parser := &attrParser{}
	expected := []byte{'{'}
	actual := parser.Trigger()
	if !bytes.Equal(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
