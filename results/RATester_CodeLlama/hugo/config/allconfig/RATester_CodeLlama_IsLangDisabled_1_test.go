package allconfig

import (
	"fmt"
	"testing"
)

func TestIsLangDisabled_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	lang := "en"
	if c.IsLangDisabled(lang) {
		t.Errorf("IsLangDisabled() = %v, want %v", c.IsLangDisabled(lang), false)
	}
}
