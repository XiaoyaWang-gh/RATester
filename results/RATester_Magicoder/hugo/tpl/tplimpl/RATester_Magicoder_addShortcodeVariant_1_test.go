package tplimpl

import (
	"fmt"
	"testing"
)

func TestaddShortcodeVariant_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	handler := &templateHandler{
		shortcodes: make(map[string]*shortcodeTemplates),
	}

	ts := &templateState{
		info: templateInfo{
			name: "test",
		},
	}

	handler.addShortcodeVariant(ts)

	shortcodename, variants := templateNameAndVariants(templateBaseName(templateShortcode, ts.Name()))

	templs, found := handler.shortcodes[shortcodename]
	if !found {
		t.Errorf("Expected shortcode to be found, but it was not")
	}

	sv := templs.indexOf(variants)

	if sv != -1 {
		if !isInternal(ts.Name()) {
			t.Errorf("Expected shortcode variant to be updated, but it was not")
		}
	} else {
		t.Errorf("Expected shortcode variant to be appended, but it was not")
	}
}
