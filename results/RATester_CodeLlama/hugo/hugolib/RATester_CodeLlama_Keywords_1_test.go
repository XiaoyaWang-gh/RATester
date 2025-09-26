package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/resources/page/pagemeta"
)

func TestKeywords_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageMeta{}
	p.pageConfig = &pagemeta.PageConfig{}
	p.pageConfig.Keywords = []string{"foo", "bar"}
	assert.Equal(t, p.Keywords(), []string{"foo", "bar"})
}
