package client

import (
	"fmt"
	"testing"
)

func TestSetMaxRedirects_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{
		maxRedirects: 3,
	}

	count := 5
	r.SetMaxRedirects(count)

	if r.maxRedirects != count {
		t.Errorf("Expected maxRedirects to be %d, but got %d", count, r.maxRedirects)
	}
}
