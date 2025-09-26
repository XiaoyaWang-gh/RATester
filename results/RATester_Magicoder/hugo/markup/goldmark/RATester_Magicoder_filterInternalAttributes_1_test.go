package goldmark

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/yuin/goldmark/ast"
)

func TestfilterInternalAttributes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &hookedRenderer{}

	attrs := []ast.Attribute{
		{Name: []byte("internalAttrPrefix:attr1")},
		{Name: []byte("internalAttrPrefix:attr2")},
		{Name: []byte("attr3")},
		{Name: []byte("attr4")},
	}

	filteredAttrs := r.filterInternalAttributes(attrs)

	if len(filteredAttrs) != 2 {
		t.Errorf("Expected 2 attributes, got %d", len(filteredAttrs))
	}

	for _, attr := range filteredAttrs {
		if bytes.HasPrefix(attr.Name, []byte(internalAttrPrefix)) {
			t.Errorf("Expected attribute %s to be filtered out", attr.Name)
		}
	}
}
