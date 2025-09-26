package allconfig

import (
	"fmt"
	"testing"
)

func TestIsMainSectionsSet_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := &ConfigCompiled{
		MainSections: []string{"section1", "section2"},
	}

	if !config.IsMainSectionsSet() {
		t.Error("Expected IsMainSectionsSet to return true, but it returned false")
	}

	config.MainSections = nil

	if config.IsMainSectionsSet() {
		t.Error("Expected IsMainSectionsSet to return false, but it returned true")
	}
}
