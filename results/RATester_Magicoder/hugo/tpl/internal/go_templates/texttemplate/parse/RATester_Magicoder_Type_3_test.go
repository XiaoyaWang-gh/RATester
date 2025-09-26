package parse

import (
	"fmt"
	"testing"
)

func TestType_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &elseNode{}
	expected := nodeElse
	actual := e.Type()
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
