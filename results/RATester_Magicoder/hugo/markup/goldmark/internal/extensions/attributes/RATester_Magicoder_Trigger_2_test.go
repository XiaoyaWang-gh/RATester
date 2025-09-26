package attributes

import (
	"fmt"
	"reflect"
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
	result := parser.Trigger()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
