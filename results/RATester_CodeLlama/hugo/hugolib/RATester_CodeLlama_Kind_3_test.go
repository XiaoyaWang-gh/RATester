package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources/page/pagemeta"
)

func TestKind_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageMeta{}
	p.pageConfig = &pagemeta.PageConfig{}
	p.pageConfig.Kind = "kind"
	if p.Kind() != "kind" {
		t.Errorf("Kind() = %v, want %v", p.Kind(), "kind")
	}
}
