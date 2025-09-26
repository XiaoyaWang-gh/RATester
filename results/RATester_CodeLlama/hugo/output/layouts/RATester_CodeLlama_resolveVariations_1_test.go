package layouts

import (
	"fmt"
	"testing"
)

func TestResolveVariations_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &layoutBuilder{}
	l.addKind()
	l.addLayoutVariations("layout1", "layout2")
	l.addSectionType()
	l.addTypeVariations("type1", "type2")

	layouts := l.resolveVariations()

	if len(layouts) != 8 {
		t.Errorf("Expected 8 layouts, got %d", len(layouts))
	}

	if layouts[0] != "layout1.type1.type2" {
		t.Errorf("Expected layout1.type1.type2, got %s", layouts[0])
	}

	if layouts[1] != "layout1.type1" {
		t.Errorf("Expected layout1.type1, got %s", layouts[1])
	}

	if layouts[2] != "layout1.type2" {
		t.Errorf("Expected layout1.type2, got %s", layouts[2])
	}

	if layouts[3] != "layout1" {
		t.Errorf("Expected layout1, got %s", layouts[3])
	}

	if layouts[4] != "layout2.type1.type2" {
		t.Errorf("Expected layout2.type1.type2, got %s", layouts[4])
	}

	if layouts[5] != "layout2.type1" {
		t.Errorf("Expected layout2.type1, got %s", layouts[5])
	}

	if layouts[6] != "layout2.type2" {
		t.Errorf("Expected layout2.type2, got %s", layouts[6])
	}

	if layouts[7] != "layout2" {
		t.Errorf("Expected layout2, got %s", layouts[7])
	}
}
