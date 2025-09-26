package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources/page/pagemeta"
)

func TestLayout_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageMeta{}
	p.pageConfig = &pagemeta.PageConfig{}
	p.pageConfig.Layout = "test"
	if p.Layout() != "test" {
		t.Errorf("Layout() = %v, want %v", p.Layout(), "test")
	}
}
