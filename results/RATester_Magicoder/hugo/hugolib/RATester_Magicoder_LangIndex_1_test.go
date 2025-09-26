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

	r := &resourceSource{
		langIndex: 1,
	}

	expected := 1
	actual := r.LangIndex()

	if actual != expected {
		t.Errorf("Expected: %d, Actual: %d", expected, actual)
	}
}
