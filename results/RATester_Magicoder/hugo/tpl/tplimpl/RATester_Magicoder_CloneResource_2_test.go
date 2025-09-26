package tplimpl

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deps"
)

func TestCloneResource_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	dst := &deps.Deps{}
	src := &deps.Deps{}

	err := dst.TemplateProvider.CloneResource(dst, src)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
