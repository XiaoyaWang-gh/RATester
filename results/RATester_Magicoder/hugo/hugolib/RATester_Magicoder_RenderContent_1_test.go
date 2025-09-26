package hugolib

import (
	"context"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestRenderContent_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	pco := &pageContentOutput{
		po: &pageOutput{
			p: &pageState{
				pid: 1,
			},
		},
		otherOutputs: maps.NewCache[uint64, *pageContentOutput](),
	}
	content := []byte("test content")
	doc := any("test doc")

	result, ok, err := pco.RenderContent(ctx, content, doc)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !ok {
		t.Errorf("Expected ok to be true")
	}
	if result == nil {
		t.Errorf("Expected result to not be nil")
	}
}
