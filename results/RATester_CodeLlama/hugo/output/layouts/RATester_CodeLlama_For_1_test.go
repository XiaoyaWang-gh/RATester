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

	l := &LayoutHandler{
		cache: map[LayoutDescriptor][]string{},
	}

	d := LayoutDescriptor{
		Type: "page",
	}

	layouts, err := l.For(d)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if len(layouts) != 0 {
		t.Errorf("expected no layouts, got %q", layouts)
	}
}
