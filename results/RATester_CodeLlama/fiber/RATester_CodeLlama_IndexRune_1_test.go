package fiber

import (
	"fmt"
	"testing"
)

func TestIndexRune_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	str := "hello world"
	needle := int32('w')
	if !IndexRune(str, needle) {
		t.Errorf("IndexRune(%q, %d) = false, want true", str, needle)
	}
}
