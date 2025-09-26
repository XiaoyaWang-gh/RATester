package fiber

import (
	"fmt"
	"testing"
)

func TestisInCharset_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	charset := []byte{'a', 'b', 'c', 'd', 'e'}

	// Test case 1: searchChar is in the charset
	searchChar := byte('b')
	if !isInCharset(searchChar, charset) {
		t.Errorf("Expected true, but got false")
	}

	// Test case 2: searchChar is not in the charset
	searchChar = byte('f')
	if isInCharset(searchChar, charset) {
		t.Errorf("Expected false, but got true")
	}
}
