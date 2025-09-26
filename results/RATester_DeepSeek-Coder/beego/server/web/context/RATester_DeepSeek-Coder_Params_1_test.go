package context

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParams_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		pnames:  []string{"name", "age"},
		pvalues: []string{"John", "30"},
	}

	expected := map[string]string{
		"name": "John",
		"age":  "30",
	}

	result := input.Params()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
