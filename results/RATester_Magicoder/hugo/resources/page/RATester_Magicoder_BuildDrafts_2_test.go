package page

import (
	"fmt"
	"testing"
)

func TestBuildDrafts_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	site := testSite{}
	result := site.BuildDrafts()
	if result != false {
		t.Errorf("Expected false, but got %v", result)
	}
}
