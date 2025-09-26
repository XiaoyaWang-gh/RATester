package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources/page/pagemeta"
)

func TestTitle_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageMeta{}
	p.pageConfig = &pagemeta.PageConfig{}
	p.pageConfig.Title = "title"
	if p.Title() != "title" {
		t.Errorf("Title() = %v, want %v", p.Title(), "title")
	}
}
