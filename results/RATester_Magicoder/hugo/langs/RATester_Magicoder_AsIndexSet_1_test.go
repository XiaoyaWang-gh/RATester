package langs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAsIndexSet_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	languages := Languages{
		{Lang: "en"},
		{Lang: "no"},
		{Lang: "de"},
	}

	expected := map[string]int{
		"en": 0,
		"no": 1,
		"de": 2,
	}

	result := languages.AsIndexSet()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
