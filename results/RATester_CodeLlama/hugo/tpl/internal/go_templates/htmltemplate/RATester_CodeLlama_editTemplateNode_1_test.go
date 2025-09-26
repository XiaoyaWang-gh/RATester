package template

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestEditTemplateNode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	e := &escaper{}
	n := &parse.TemplateNode{}
	callee := "callee"

	// When
	e.editTemplateNode(n, callee)

	// Then
	assert.Equal(t, callee, e.templateNodeEdits[n])
}
