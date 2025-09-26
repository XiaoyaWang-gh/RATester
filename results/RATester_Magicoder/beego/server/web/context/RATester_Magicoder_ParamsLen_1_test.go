package context

import (
	"fmt"
	"testing"
)

func TestParamsLen_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	input := &BeegoInput{
		pnames: []string{"param1", "param2", "param3"},
	}

	expected := 3
	actual := input.ParamsLen()

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
