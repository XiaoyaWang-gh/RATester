package page

import (
	"fmt"
	"testing"
)

func TestRSSLink_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	nop := nopPage(0)
	url := nop.RSSLink()

	if url != "" {
		t.Errorf("Expected empty URL, got %v", url)
	}
}
