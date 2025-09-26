package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources/page/pagemeta"
)

func TestNoRender_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageMeta{}
	p.pageConfig.Build.Render = pagemeta.Always
	if p.noRender() {
		t.Error("expected noRender to be false")
	}
	p.pageConfig.Build.Render = pagemeta.Never
	if !p.noRender() {
		t.Error("expected noRender to be true")
	}
}
