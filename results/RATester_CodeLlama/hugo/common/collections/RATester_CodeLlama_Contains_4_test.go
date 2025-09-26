package collections

import (
	"fmt"
	"testing"
)

func TestContains_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ss := SortedStringSlice{"a", "b", "c"}
	if !ss.Contains("a") {
		t.Errorf("ss.Contains(\"a\") = false, want true")
	}
	if ss.Contains("d") {
		t.Errorf("ss.Contains(\"d\") = true, want false")
	}
}
