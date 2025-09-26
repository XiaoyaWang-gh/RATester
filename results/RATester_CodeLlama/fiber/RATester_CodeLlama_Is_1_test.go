package fiber

import (
	"fmt"
	"testing"
)

func TestIs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var c *DefaultCtx
	var extension string
	var want bool
	var got bool

	// TEST CASE 1:
	// Test case with extension is empty.
	extension = ""
	want = false
	got = c.Is(extension)
	if got != want {
		t.Errorf("c.Is(extension) = %v, want %v", got, want)
	}

	// TEST CASE 2:
	// Test case with extension is not empty.
	extension = "json"
	want = true
	got = c.Is(extension)
	if got != want {
		t.Errorf("c.Is(extension) = %v, want %v", got, want)
	}
}
