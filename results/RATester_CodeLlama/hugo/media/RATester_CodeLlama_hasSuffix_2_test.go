package media

import (
	"fmt"
	"testing"
)

func TestHasSuffix_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Type{
		SuffixesCSV: "jpg,jpeg",
	}
	if !m.hasSuffix("jpeg") {
		t.Errorf("expected true, got false")
	}
	if m.hasSuffix("png") {
		t.Errorf("expected false, got true")
	}
}
