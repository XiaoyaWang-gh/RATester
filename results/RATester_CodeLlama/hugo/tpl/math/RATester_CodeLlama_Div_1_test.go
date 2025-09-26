package math

import (
	"fmt"
	"testing"
)

func TestDiv_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}
	inputs := []any{10, 2}
	result, err := ns.Div(inputs...)
	if err != nil {
		t.Fatal(err)
	}
	if result != 5 {
		t.Errorf("expected 5, got %v", result)
	}
}
