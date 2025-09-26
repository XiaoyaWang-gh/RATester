package allconfig

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetMainSections_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cc := &ConfigCompiled{
		MainSections: []string{"section1", "section2"},
	}

	newSections := []string{"section3", "section4"}
	cc.SetMainSections(newSections)

	if !reflect.DeepEqual(cc.MainSections, newSections) {
		t.Errorf("Expected MainSections to be %v, but got %v", newSections, cc.MainSections)
	}
}
