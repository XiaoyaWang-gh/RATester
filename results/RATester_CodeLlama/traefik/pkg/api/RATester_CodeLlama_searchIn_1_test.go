package api

import (
	"fmt"
	"testing"
)

func TestSearchIn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &searchCriterion{
		Search: "test",
	}

	values := []string{"test", "test1", "test2"}

	if !c.searchIn(values...) {
		t.Errorf("searchIn() = %v, want %v", false, true)
	}
}
