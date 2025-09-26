package binder

import (
	"fmt"
	"testing"
)

func TestName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &xmlBinding{}
	name := b.Name()
	if name != "xml" {
		t.Errorf("Expected 'xml', but got '%s'", name)
	}
}
