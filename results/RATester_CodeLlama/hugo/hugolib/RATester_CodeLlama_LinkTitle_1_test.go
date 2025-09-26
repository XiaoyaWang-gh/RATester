package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources/page/pagemeta"
)

func TestLinkTitle_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageMeta{}
	p.pageConfig = &pagemeta.PageConfig{}
	p.pageConfig.LinkTitle = "link title"
	if p.LinkTitle() != "link title" {
		t.Errorf("LinkTitle() = %v, want %v", p.LinkTitle(), "link title")
	}
}
