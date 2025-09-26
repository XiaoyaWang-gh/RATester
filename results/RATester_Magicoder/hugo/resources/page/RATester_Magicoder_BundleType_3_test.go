package page

import (
	"fmt"
	"testing"
)

func TestBundleType_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	expected := ""
	result := nop.BundleType()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
