package page

import (
	"fmt"
	"testing"
)

func TestURL_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	url := nop.URL()
	if url != "" {
		t.Errorf("Expected URL to be empty, but got %s", url)
	}
}
