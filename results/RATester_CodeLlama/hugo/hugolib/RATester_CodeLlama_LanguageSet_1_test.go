package hugolib

import (
	"fmt"
	"testing"
)

func TestLanguageSet_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var h HugoSites
	set := h.LanguageSet()
	if len(set) != 0 {
		t.Errorf("LanguageSet() = %v, want %v", set, 0)
	}
}
