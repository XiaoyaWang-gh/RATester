package page

import (
	"fmt"
	"testing"
)

func TestLastmod_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	got := nop.Lastmod()
	if !got.IsZero() {
		t.Errorf("Expected Lastmod() to return a zero time, but got: %v", got)
	}
}
