package layouts

import (
	"fmt"
	"testing"
)

func TestaddTypeVariations_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &layoutBuilder{
		layoutVariations: []string{},
		typeVariations:   []string{},
		d: LayoutDescriptor{
			RenderingHook: true,
		},
	}

	l.addTypeVariations("test1", "test2")

	if len(l.typeVariations) != 2 {
		t.Errorf("Expected 2 type variations, got %d", len(l.typeVariations))
	}

	if l.typeVariations[0] != "test1"+renderingHookRoot {
		t.Errorf("Expected first type variation to be 'test1' + renderingHookRoot, got %s", l.typeVariations[0])
	}

	if l.typeVariations[1] != "test2"+renderingHookRoot {
		t.Errorf("Expected second type variation to be 'test2' + renderingHookRoot, got %s", l.typeVariations[1])
	}
}
