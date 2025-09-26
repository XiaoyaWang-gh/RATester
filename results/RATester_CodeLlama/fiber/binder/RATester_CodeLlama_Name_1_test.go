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
	if b.Name() != "xml" {
		t.Errorf("Expected %s, got %s", "xml", b.Name())
	}
}
