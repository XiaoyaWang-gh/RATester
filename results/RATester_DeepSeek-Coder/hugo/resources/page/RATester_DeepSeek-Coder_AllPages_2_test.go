package page

import (
	"fmt"
	"testing"
)

func TestAllPages_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	site := testSite{}
	pages := site.AllPages()
	if pages != nil {
		t.Errorf("Expected nil, got %v", pages)
	}
}
