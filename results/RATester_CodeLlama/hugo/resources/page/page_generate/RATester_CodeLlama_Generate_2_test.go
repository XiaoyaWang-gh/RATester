package page_generate

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/codegen"
)

func TestGenerate_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &codegen.Inspector{}
	if err := Generate(c); err != nil {
		t.Errorf("failed to generate JSON marshaler: %v", err)
	}
}
