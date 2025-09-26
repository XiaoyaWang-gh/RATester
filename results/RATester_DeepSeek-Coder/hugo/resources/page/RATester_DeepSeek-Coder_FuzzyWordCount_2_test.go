package page

import (
	"context"
	"fmt"
	"testing"
)

func TestFuzzyWordCount_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	n := new(nopContent)
	wordCount := n.FuzzyWordCount(ctx)
	if wordCount != 0 {
		t.Errorf("Expected word count to be 0, got %d", wordCount)
	}
}
