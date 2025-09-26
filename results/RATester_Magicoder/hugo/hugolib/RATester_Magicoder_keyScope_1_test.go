package hugolib

import (
	"context"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/hugo"
	"github.com/gohugoio/hugo/output"
)

func TestkeyScope_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	ccs := &cachedContentScope{
		pco: &pageContentOutput{
			po: &pageOutput{
				f: output.Format{},
			},
		},
	}
	expected := hugo.GetMarkupScope(ctx) + ccs.pco.po.f.Name
	result := ccs.keyScope(ctx)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
