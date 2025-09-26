package page

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/langs"
)

func TestLanguageCode_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testSite := testSite{
		l: &langs.Language{Lang: "en"},
	}

	expected := "en"
	actual := testSite.LanguageCode()

	if actual != expected {
		t.Errorf("Expected LanguageCode to be %s, but got %s", expected, actual)
	}
}
