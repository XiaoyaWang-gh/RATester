package csrf

import (
	"fmt"
	"testing"
)

func TestMatch_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := subdomain{prefix: "a", suffix: "b"}
	o := "ab"
	if !s.match(o) {
		t.Errorf("s.match(o) = %v, want %v", false, true)
	}
}
