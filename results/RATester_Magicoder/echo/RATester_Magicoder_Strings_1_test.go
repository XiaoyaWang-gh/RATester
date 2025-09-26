package echo

import (
	"fmt"
	"testing"
)

func TestStrings_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{
		ValuesFunc: func(sourceParam string) []string {
			if sourceParam == "test" {
				return []string{"value1", "value2"}
			}
			return nil
		},
	}

	var dest []string
	b.Strings("test", &dest)

	if len(dest) != 2 || dest[0] != "value1" || dest[1] != "value2" {
		t.Errorf("Expected dest to be [value1, value2], but got %v", dest)
	}
}
