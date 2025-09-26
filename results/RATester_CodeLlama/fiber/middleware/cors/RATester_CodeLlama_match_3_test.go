package cors

import (
	"fmt"
	"testing"
)

func TestMatch_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := subdomain{
		prefix: "*.",
		suffix: ".com",
	}
	o := "example.com"
	if !s.match(o) {
		t.Errorf("Expected %v to match %v", s, o)
	}
}
