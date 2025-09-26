package fiber

import (
	"fmt"
	"testing"
)

func TestFindNextCharsetPositionConstraint_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	search := "abcdefghijklmnopqrstuvwxyz"
	charset := []byte("abcdefghijklmnopqrstuvwxyz")
	expected := 1
	actual := findNextCharsetPositionConstraint(search, charset)
	if actual != expected {
		t.Errorf("Expected %d, actual %d", expected, actual)
	}
}
