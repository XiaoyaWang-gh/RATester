package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/config/allconfig"
	"github.com/gohugoio/hugo/langs"
)

func TestLanguages_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Site{
		h: &HugoSites{
			Configs: &allconfig.Configs{
				Languages: langs.Languages{
					{Lang: "en"},
					{Lang: "fr"},
				},
			},
		},
	}

	languages := s.Languages()

	if len(languages) != 2 {
		t.Errorf("Expected 2 languages, got %d", len(languages))
	}

	if languages[0].Lang != "en" {
		t.Errorf("Expected language 'en', got '%s'", languages[0].Lang)
	}

	if languages[1].Lang != "fr" {
		t.Errorf("Expected language 'fr', got '%s'", languages[1].Lang)
	}
}
