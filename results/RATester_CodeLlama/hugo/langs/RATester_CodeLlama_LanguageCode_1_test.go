package langs

import (
	"fmt"
	"testing"
)

func TestLanguageCode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &Language{
		Lang: "en",
	}
	if l.LanguageCode() != "en" {
		t.Errorf("LanguageCode() = %q, want %q", l.LanguageCode(), "en")
	}
}
