package tplimpl

import (
	"fmt"
	"testing"
)

func TestIndexOf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	variants := []string{"a", "b", "c"}
	s := &shortcodeTemplates{
		variants: []shortcodeVariant{
			{
				variants: []string{"a", "b", "c"},
			},
			{
				variants: []string{"a", "b", "c"},
			},
			{
				variants: []string{"a", "b", "c"},
			},
		},
	}
	if s.indexOf(variants) != 0 {
		t.Errorf("expected 0, got %d", s.indexOf(variants))
	}
}
