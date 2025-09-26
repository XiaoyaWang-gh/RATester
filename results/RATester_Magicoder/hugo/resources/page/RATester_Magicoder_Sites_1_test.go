package page

import (
	"fmt"
	"testing"
)

func TestSites_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.Sites() != nil {
		t.Errorf("Expected Sites() to return nil, but got %v", p.Sites())
	}
}
