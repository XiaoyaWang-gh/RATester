package images

import (
	"fmt"
	"testing"
)

func TestString_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := resamp{name: "TestName"}
	expected := "TestName"
	if r.String() != expected {
		t.Errorf("Expected %s, got %s", expected, r.String())
	}
}
