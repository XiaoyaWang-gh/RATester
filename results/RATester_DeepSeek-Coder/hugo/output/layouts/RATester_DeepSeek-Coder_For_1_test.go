package layouts

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFor_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	l := &LayoutHandler{
		cache: make(map[LayoutDescriptor][]string),
	}

	d := LayoutDescriptor{
		Type:             "type",
		Section:          "section",
		Kind:             "kind",
		KindVariants:     "variants",
		Lang:             "lang",
		Layout:           "layout",
		OutputFormatName: "format",
		Suffix:           "suffix",
	}

	layouts := []string{"layout1", "layout2"}

	l.cache[d] = layouts

	result, err := l.For(d)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !reflect.DeepEqual(result, layouts) {
		t.Errorf("Expected %v, got %v", layouts, result)
	}
}
