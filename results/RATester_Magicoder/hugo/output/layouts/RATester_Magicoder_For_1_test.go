package layouts

import (
	"fmt"
	"testing"
)

func TestFor_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lh := &LayoutHandler{
		cache: make(map[LayoutDescriptor][]string),
	}

	d := LayoutDescriptor{
		Type:    "page",
		Section: "section",
		Kind:    "kind",
		Lang:    "en",
		Layout:  "layout",
	}

	layouts, err := lh.For(d)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(layouts) != 0 {
		t.Errorf("Expected 0 layouts, got %d", len(layouts))
	}

	lh.cache[d] = []string{"layout1", "layout2"}

	layouts, err = lh.For(d)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(layouts) != 2 {
		t.Errorf("Expected 2 layouts, got %d", len(layouts))
	}
}
