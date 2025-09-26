package media

import (
	"fmt"
	"testing"
)

func TestIsHTML_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Type{
		SubType: "html",
	}
	if !m.IsHTML() {
		t.Errorf("IsHTML() = %v, want %v", false, true)
	}
}
