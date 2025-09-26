package hugolib

import (
	"context"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestParseAndRenderContent_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	pageContentOutput := &pageContentOutput{
		po: &pageOutput{
			p: &pageState{
				pid: 1,
			},
		},
		otherOutputs: maps.NewCache[uint64, *pageContentOutput](),
	}
	content := []byte("test content")
	renderTOC := true
	result, err := pageContentOutput.ParseAndRenderContent(ctx, content, renderTOC)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("Expected a result, but got nil")
	}
}
