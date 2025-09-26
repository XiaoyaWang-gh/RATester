package template

import (
	"fmt"
	"testing"
)

func TestIndexTagEnd_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := []byte("</a>")
	tag := []byte("</a>")
	if indexTagEnd(s, tag) != 0 {
		t.Errorf("indexTagEnd(%v, %v) = %v; want 0", s, tag, indexTagEnd(s, tag))
	}
}
