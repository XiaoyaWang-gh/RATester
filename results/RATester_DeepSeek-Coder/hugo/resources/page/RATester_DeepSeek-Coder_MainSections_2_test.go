package page

import (
	"fmt"
	"testing"
)

func TestMainSections_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	site := testSite{}
	sections := site.MainSections()
	if sections != nil {
		t.Errorf("Expected nil, got %v", sections)
	}
}
