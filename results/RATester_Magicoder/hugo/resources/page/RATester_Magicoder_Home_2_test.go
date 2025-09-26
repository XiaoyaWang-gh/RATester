package page

import (
	"fmt"
	"testing"
)

func TestHome_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	site := testSite{}
	page := site.Home()
	if page != nil {
		t.Errorf("Expected nil, got %v", page)
	}
}
