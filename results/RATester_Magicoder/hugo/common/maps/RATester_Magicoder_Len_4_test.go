package maps

import (
	"fmt"
	"testing"
)

func TestLen_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := &Cache[int, string]{m: map[int]string{1: "one", 2: "two", 3: "three"}}
	expected := 3
	actual := cache.Len()
	if actual != expected {
		t.Errorf("Expected Len() to return %d, but got %d", expected, actual)
	}
}
