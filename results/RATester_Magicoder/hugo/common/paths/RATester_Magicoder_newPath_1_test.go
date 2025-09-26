package paths

import (
	"fmt"
	"testing"
)

func TestnewPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pp := &PathParser{}
	path := pp.newPath("component")

	if path.component != "component" {
		t.Errorf("Expected component to be 'component', but got %s", path.component)
	}

	if path.posContainerLow != -1 {
		t.Errorf("Expected posContainerLow to be -1, but got %d", path.posContainerLow)
	}

	if path.posContainerHigh != -1 {
		t.Errorf("Expected posContainerHigh to be -1, but got %d", path.posContainerHigh)
	}

	if path.posSectionHigh != -1 {
		t.Errorf("Expected posSectionHigh to be -1, but got %d", path.posSectionHigh)
	}

	if path.posIdentifierLanguage != -1 {
		t.Errorf("Expected posIdentifierLanguage to be -1, but got %d", path.posIdentifierLanguage)
	}
}
