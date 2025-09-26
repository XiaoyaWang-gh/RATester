package page

import (
	"fmt"
	"testing"
)

func TestAllPages_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	site := testSite{}
	pages := site.AllPages()
	if pages != nil {
		t.Errorf("Expected AllPages() to return nil, but got %v", pages)
	}
}
