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

	c := &ConfigCompiled{}
	if c.IsMainSectionsSet() {
		t.Error("IsMainSectionsSet() should return false when MainSections is nil")
	}
	c.SetMainSections([]string{"a"})
	if !c.IsMainSectionsSet() {
		t.Error("IsMainSectionsSet() should return true when MainSections is set")
	}
}
