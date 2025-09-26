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
	expected := "header"
	actual := h.Name()
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
