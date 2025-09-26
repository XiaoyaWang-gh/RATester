package paths

import (
	"fmt"
	"testing"
)

func TestLang_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Paths{}
	if got := p.Lang(); got != "" {
		t.Errorf("Lang() = %v, want %v", got, "")
	}
}
