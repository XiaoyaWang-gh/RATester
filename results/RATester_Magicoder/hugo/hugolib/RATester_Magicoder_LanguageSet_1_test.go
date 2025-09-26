package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/langs"
)

func TestLanguageSet_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hugoSites := &HugoSites{
		Sites: []*Site{
			{
				language: &langs.Language{
					Lang: "en",
				},
			},
			{
				language: &langs.Language{
					Lang: "fr",
				},
			},
		},
	}

	expected := map[string]int{
		"en": 0,
		"fr": 1,
	}

	result := hugoSites.LanguageSet()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
