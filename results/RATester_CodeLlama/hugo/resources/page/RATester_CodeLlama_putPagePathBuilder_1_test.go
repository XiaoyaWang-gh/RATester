package page

import (
	"fmt"
	"testing"
)

func TestPutPagePathBuilder_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := pagePathBuilderPool.Get().(*pagePathBuilder)
	b.els = []string{"a", "b", "c"}
	b.fullSuffix = "d"
	b.baseNameSameAsType = true
	b.isUgly = true
	b.noSubResources = true
	b.prefixLink = "e"
	b.prefixPath = "f"
	b.linkUpperOffset = 1
	putPagePathBuilder(b)
	if b.els != nil {
		t.Errorf("b.els = %v; want nil", b.els)
	}
	if b.fullSuffix != "" {
		t.Errorf("b.fullSuffix = %q; want \"\"", b.fullSuffix)
	}
	if b.baseNameSameAsType != false {
		t.Errorf("b.baseNameSameAsType = %v; want false", b.baseNameSameAsType)
	}
	if b.isUgly != false {
		t.Errorf("b.isUgly = %v; want false", b.isUgly)
	}
	if b.noSubResources != false {
		t.Errorf("b.noSubResources = %v; want false", b.noSubResources)
	}
	if b.prefixLink != "" {
		t.Errorf("b.prefixLink = %q; want \"\"", b.prefixLink)
	}
	if b.prefixPath != "" {
		t.Errorf("b.prefixPath = %q; want \"\"", b.prefixPath)
	}
	if b.linkUpperOffset != 0 {
		t.Errorf("b.linkUpperOffset = %d; want 0", b.linkUpperOffset)
	}
}
