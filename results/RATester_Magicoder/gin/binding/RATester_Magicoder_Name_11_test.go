package binding

import (
	"fmt"
	"testing"
)

func TestName_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := headerBinding{}
	name := h.Name()
	if name != "header" {
		t.Errorf("Expected 'header', but got '%s'", name)
	}
}
