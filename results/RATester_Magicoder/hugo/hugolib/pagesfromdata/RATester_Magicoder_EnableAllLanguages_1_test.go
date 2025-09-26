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

	ctx := &pagesFromDataTemplateContext{
		p: &PagesFromTemplate{
			buildState: &BuildState{},
		},
	}

	ctx.EnableAllLanguages()

	if !ctx.p.buildState.EnableAllLanguages {
		t.Error("EnableAllLanguages() did not set EnableAllLanguages to true")
	}
}
