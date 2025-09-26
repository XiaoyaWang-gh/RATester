package tcp

import (
	"fmt"
	"testing"
)

func TestThen_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := Chain{}
	h := HandlerFunc(func(conn WriteCloser) {})
	h2, err := c.Then(h)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if h2 == nil {
		t.Errorf("expected a handler, got nil")
	}
}
