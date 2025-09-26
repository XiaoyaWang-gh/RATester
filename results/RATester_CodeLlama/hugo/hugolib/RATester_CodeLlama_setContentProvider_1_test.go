package hugolib

import (
	"fmt"
	"testing"
)

func TestSetContentProvider_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageOutput{}
	cp := &pageContentOutput{}
	p.setContentProvider(cp)
	if p.contentRenderer != cp {
		t.Errorf("contentRenderer not set")
	}
	if p.ContentProvider != cp {
		t.Errorf("ContentProvider not set")
	}
	if p.MarkupProvider != cp {
		t.Errorf("MarkupProvider not set")
	}
	if p.PageRenderProvider != cp {
		t.Errorf("PageRenderProvider not set")
	}
	if p.TableOfContentsProvider != cp {
		t.Errorf("TableOfContentsProvider not set")
	}
	if p.RenderShortcodesProvider != cp {
		t.Errorf("RenderShortcodesProvider not set")
	}
	if p.pco != cp {
		t.Errorf("pco not set")
	}
}
