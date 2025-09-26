package validation

import (
	"fmt"
	"testing"
)

func TestIsSatisfied_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := Enum{
		Rules: "a|b|c",
		Key:   "key",
	}
	i := "a"
	if !e.IsSatisfied(i) {
		t.Errorf("IsSatisfied() = %v, want %v", false, true)
	}
}
