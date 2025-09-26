package validation

import (
	"fmt"
	"testing"
)

func TestDefaultMessage_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := Range{
		Min: Min{Min: 1},
		Max: Max{Max: 10},
		Key: "test",
	}
	expected := fmt.Sprintf(MessageTmpls["Range"], 1, 10)
	actual := r.DefaultMessage()
	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
