package source

import (
	"fmt"
	"testing"
)

func TestTranslationBaseName_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fi := &File{}
	if got := fi.TranslationBaseName(); got != "" {
		t.Errorf("TranslationBaseName() = %v, want %v", got, "")
	}
}
