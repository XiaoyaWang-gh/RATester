package binding

import (
	"fmt"
	"testing"
)

func TestName_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := xmlBinding{}
	if b.Name() != "xml" {
		t.Errorf("Expected 'xml', got '%s'", b.Name())
	}
}
