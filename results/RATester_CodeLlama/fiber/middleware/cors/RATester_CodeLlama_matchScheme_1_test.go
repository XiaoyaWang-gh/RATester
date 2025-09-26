package cors

import (
	"fmt"
	"testing"
)

func TestMatchScheme_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	domain := "example.com"
	pattern := "example.com"
	if matchScheme(domain, pattern) != true {
		t.Errorf("matchScheme(%s, %s) = %v, want %v", domain, pattern, false, true)
	}
}
