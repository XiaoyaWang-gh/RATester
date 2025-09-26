package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources/page/pagemeta"
)

func TestWeight_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageMeta{}
	p.pageConfig = &pagemeta.PageConfig{Weight: 10}
	if p.Weight() != 10 {
		t.Errorf("expected 10, got %d", p.Weight())
	}
}
