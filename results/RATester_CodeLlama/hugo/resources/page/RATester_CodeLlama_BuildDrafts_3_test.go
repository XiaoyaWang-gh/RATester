package page

import (
	"fmt"
	"testing"
)

func TestBuildDrafts_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := testSite{}
	if s.BuildDrafts() != false {
		t.Errorf("BuildDrafts() = %v, want %v", s.BuildDrafts(), false)
	}
}
