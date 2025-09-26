package hexec

import (
	"fmt"
	"testing"
)

func TestError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &NotFoundError{
		name:   "test",
		method: "test",
	}

	expected := fmt.Sprintf("binary with name %q not found %s", e.name, e.method)
	actual := e.Error()

	if actual != expected {
		t.Errorf("Expected %q, got %q", expected, actual)
	}
}
