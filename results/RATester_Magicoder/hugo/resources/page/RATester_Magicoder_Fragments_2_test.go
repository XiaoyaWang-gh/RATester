package page

import (
	"context"
	"fmt"
	"testing"
)

func TestFragments_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	n := nopMarkup(0)
	result := n.Fragments(ctx)
	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}
