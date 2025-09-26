package page

import (
	"fmt"
	"testing"
)

func TestSocial_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	site := testSite{}
	result := site.Social()
	if result == nil {
		t.Error("Expected a non-nil map, but got nil")
	}
}
