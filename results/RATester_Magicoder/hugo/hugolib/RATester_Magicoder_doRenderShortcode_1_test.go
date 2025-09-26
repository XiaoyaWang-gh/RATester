package hugolib

import (
	"context"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/tpl"
)

func TestdoRenderShortcode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	level := 0
	s := &Site{}
	tplVariants := tpl.TemplateVariants{}
	sc := &shortcode{}
	parent := &ShortcodeWithPage{}
	p := &pageState{}
	isRenderString := false

	_, err := doRenderShortcode(ctx, level, s, tplVariants, sc, parent, p, isRenderString)
	if err != nil {
		t.Errorf("doRenderShortcode returned an error: %v", err)
	}
}
