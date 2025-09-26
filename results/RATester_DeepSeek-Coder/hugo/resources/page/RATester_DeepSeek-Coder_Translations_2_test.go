package page

import (
	"fmt"
	"testing"
)

func TestTranslations_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	translations := p.Translations()
	if translations != nil {
		t.Errorf("Expected nil, got %v", translations)
	}
}
