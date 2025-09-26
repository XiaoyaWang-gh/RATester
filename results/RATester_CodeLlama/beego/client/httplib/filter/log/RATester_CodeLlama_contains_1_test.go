package log

import (
	"fmt"
	"testing"
)

func TestContains_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := []string{"a", "b", "c"}
	e := "b"
	if contains(s, e) != true {
		t.Errorf("contains(%v, %v) = %v; want true", s, e, contains(s, e))
	}
}
