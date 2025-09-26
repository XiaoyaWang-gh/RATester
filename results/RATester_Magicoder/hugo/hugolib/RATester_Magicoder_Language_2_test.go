package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/langs"
)

func TestLanguage_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := Site{language: &langs.Language{}}

	result := s.Language()

	if result != s.language {
		t.Errorf("Expected %v, got %v", s.language, result)
	}
}
