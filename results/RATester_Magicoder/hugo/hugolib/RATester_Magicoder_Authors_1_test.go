package hugolib

import (
	"fmt"
	"testing"
)

func TestAuthors_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	page := &pageMeta{
		term: "test",
	}

	result := page.Authors()

	if result != nil {
		t.Errorf("Expected nil, but got %v", result)
	}
}
