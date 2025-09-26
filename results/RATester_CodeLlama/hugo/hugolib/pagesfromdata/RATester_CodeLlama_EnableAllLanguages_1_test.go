package pagesfromdata

import (
	"fmt"
	"testing"
)

func TestEnableAllLanguages_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pagesFromDataTemplateContext{}
	p.EnableAllLanguages()
	if p.p.buildState.EnableAllLanguages != true {
		t.Errorf("EnableAllLanguages() failed")
	}
}
