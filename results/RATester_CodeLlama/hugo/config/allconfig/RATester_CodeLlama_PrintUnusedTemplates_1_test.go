package allconfig

import (
	"fmt"
	"testing"
)

func TestPrintUnusedTemplates_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := ConfigLanguage{}
	if c.PrintUnusedTemplates() {
		t.Error("PrintUnusedTemplates() should be false")
	}
}
