package page

import (
	"context"
	"fmt"
	"testing"
)

func TestPlain_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	n := new(nopContent)
	result := n.Plain(ctx)
	if result != "" {
		t.Errorf("Expected empty string, got %s", result)
	}
}
