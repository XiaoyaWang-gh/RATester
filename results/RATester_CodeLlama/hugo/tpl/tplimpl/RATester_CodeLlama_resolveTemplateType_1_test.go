package tplimpl

import (
	"fmt"
	"testing"
)

func TestResolveTemplateType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var name string
	var expected templateType
	var actual templateType

	// CONTEXT_0
	name = "shortcode"
	expected = templateShortcode
	actual = resolveTemplateType(name)
	if actual != expected {
		t.Errorf("resolveTemplateType(%s) = %d, expected %d", name, actual, expected)
	}

	// CONTEXT_1
	name = "partials/partial"
	expected = templatePartial
	actual = resolveTemplateType(name)
	if actual != expected {
		t.Errorf("resolveTemplateType(%s) = %d, expected %d", name, actual, expected)
	}

	// CONTEXT_2
	name = "undefined"
	expected = templateUndefined
	actual = resolveTemplateType(name)
	if actual != expected {
		t.Errorf("resolveTemplateType(%s) = %d, expected %d", name, actual, expected)
	}
}
