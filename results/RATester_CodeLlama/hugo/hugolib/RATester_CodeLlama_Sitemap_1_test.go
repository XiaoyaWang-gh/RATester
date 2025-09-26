package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/config"
	"github.com/gohugoio/hugo/resources/page/pagemeta"
)

func TestSitemap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageMeta{}
	p.pageConfig = &pagemeta.PageConfig{}
	p.pageConfig.Sitemap = config.SitemapConfig{
		Priority: 0.5,
	}
	if got := p.Sitemap(); got != p.pageConfig.Sitemap {
		t.Errorf("Sitemap() = %v, want %v", got, p.pageConfig.Sitemap)
	}
}
