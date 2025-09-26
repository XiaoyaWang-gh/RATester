package template

import (
	"fmt"
	"testing"
)

func TestIsComment_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var s state
	if got := isComment(s); got != false {
		t.Errorf("isComment(%v) = %v; want false", s, got)
	}
}
