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
	expected := "xml"
	actual := b.Name()
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
