package hugolib

import (
	"fmt"
	"testing"
)

func TestTranslationKey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageState{}
	p.m.pageConfig.TranslationKey = "test"
	if p.TranslationKey() != "test" {
		t.Errorf("TranslationKey() = %q, want %q", p.TranslationKey(), "test")
	}
	p.m.pageConfig.TranslationKey = ""
	if p.TranslationKey() != p.Path() {
		t.Errorf("TranslationKey() = %q, want %q", p.TranslationKey(), p.Path())
	}
}
