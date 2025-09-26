package page

import (
	"fmt"
	"testing"
)

func TestCopyright_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	site := testSite{}
	expected := ""
	result := site.Copyright()
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
