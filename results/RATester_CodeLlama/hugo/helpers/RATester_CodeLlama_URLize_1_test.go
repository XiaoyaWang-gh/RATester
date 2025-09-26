package helpers

import (
	"fmt"
	"testing"
)

func TestURLize_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &PathSpec{}
	uri := "https://example.com/foo/bar"
	if got := p.URLize(uri); got != uri {
		t.Errorf("URLize() = %v, want %v", got, uri)
	}
}
