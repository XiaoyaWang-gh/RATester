package hugolib

import (
	"fmt"
	"testing"
)

func TestsetContentProvider_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	po := &pageOutput{}
	pco := &pageContentOutput{}

	po.setContentProvider(pco)

	if po.contentRenderer != pco {
		t.Errorf("Expected contentRenderer to be set to pco, but got %v", po.contentRenderer)
	}
	if po.ContentProvider != pco {
		t.Errorf("Expected ContentProvider to be set to pco, but got %v", po.ContentProvider)
	}
	if po.MarkupProvider != pco {
		t.Errorf("Expected MarkupProvider to be set to pco, but got %v", po.MarkupProvider)
	}
	if po.PageRenderProvider != pco {
		t.Errorf("Expected PageRenderProvider to be set to pco, but got %v", po.PageRenderProvider)
	}
	if po.TableOfContentsProvider != pco {
		t.Errorf("Expected TableOfContentsProvider to be set to pco, but got %v", po.TableOfContentsProvider)
	}
	if po.RenderShortcodesProvider != pco {
		t.Errorf("Expected RenderShortcodesProvider to be set to pco, but got %v", po.RenderShortcodesProvider)
	}
	if po.pco != pco {
		t.Errorf("Expected pco to be set to pco, but got %v", po.pco)
	}
}
