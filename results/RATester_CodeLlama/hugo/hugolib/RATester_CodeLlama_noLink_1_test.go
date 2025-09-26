package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources/page/pagemeta"
)

func TestNoLink_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageMeta{}
	p.pageConfig.Build.Render = pagemeta.Never
	if !p.noLink() {
		t.Errorf("noLink() = %v, want %v", p.noLink(), true)
	}
}
