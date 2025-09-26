package hugolib

import (
	"fmt"
	"testing"
)

func TestLangIndex_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &resourceSource{}
	if r.LangIndex() != 0 {
		t.Errorf("LangIndex() = %d, want %d", r.LangIndex(), 0)
	}
}
