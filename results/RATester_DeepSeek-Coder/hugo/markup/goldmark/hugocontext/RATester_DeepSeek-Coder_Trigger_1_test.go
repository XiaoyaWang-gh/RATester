package hugocontext

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTrigger_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &hugoContextParser{}
	expected := []byte{'{'}
	result := h.Trigger()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
