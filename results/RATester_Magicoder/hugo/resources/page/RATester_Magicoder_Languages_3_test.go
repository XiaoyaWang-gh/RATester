package page

import (
	"fmt"
	"testing"
)

func TestLanguages_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	site := testSite{}
	languages := site.Languages()
	if languages != nil {
		t.Errorf("Expected nil, but got %v", languages)
	}
}
