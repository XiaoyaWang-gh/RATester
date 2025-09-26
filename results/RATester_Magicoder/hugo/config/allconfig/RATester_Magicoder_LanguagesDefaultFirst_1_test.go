package allconfig

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/langs"
)

func TestLanguagesDefaultFirst_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cfg := &ConfigLanguage{
		m: &Configs{
			LanguagesDefaultFirst: langs.Languages{
				&langs.Language{
					Lang: "en",
				},
			},
		},
	}

	expected := langs.Languages{
		&langs.Language{
			Lang: "en",
		},
	}

	result := cfg.LanguagesDefaultFirst()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
