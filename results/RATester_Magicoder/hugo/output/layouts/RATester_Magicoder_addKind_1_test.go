package layouts

import (
	"fmt"
	"testing"
)

func TestaddKind_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &layoutBuilder{
		d: LayoutDescriptor{
			Kind: "testKind",
		},
	}

	l.addKind()

	if len(l.layoutVariations) == 0 || len(l.typeVariations) == 0 {
		t.Errorf("Expected layoutVariations and typeVariations to be populated, but they were not.")
	}
}
