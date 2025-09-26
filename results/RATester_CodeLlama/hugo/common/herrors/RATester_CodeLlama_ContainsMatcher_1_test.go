package herrors

import (
	"fmt"
	"testing"
)

func TestContainsMatcher_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	text := "foo"
	m := LineMatcher{
		Line: "foo bar",
	}
	if got := ContainsMatcher(text)(m); got != 1 {
		t.Errorf("ContainsMatcher() = %v, want %v", got, 1)
	}
}
