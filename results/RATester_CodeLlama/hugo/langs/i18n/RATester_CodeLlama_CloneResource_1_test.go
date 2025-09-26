package i18n

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestCloneResource_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var dst, src *deps.Deps
	var tp *TranslationProvider
	var err error

	dst = &deps.Deps{}
	src = &deps.Deps{}
	tp = &TranslationProvider{}

	err = tp.CloneResource(dst, src)
	if err != nil {
		t.Errorf("CloneResource() error = %v", err)
		return
	}

	if dst.Translate == nil {
		t.Errorf("CloneResource() dst.Translate = nil")
	}
}
