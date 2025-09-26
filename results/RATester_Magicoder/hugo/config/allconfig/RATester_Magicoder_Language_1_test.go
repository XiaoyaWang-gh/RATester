package allconfig

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/langs"
)

func TestLanguage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cfg := ConfigLanguage{
		language: &langs.Language{
			Lang: "en",
		},
	}

	expected := &langs.Language{
		Lang: "en",
	}

	result := cfg.Language()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %v, Got: %v", expected, result)
	}
}
