package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/page"
)

func TestTranslations_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pageState := &pageState{
		// Initialize fields here
	}

	expected := page.Pages{
		// Initialize expected pages here
	}

	result := pageState.Translations()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected: %v, Got: %v", expected, result)
	}
}
