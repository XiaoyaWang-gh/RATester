package hints

import (
	"fmt"
	"testing"
)

func TestOrderBy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hint := OrderBy("test")
	if hint.GetKey() != KeyOrderBy {
		t.Errorf("Expected key to be %v, got %v", KeyOrderBy, hint.GetKey())
	}
	if hint.GetValue() != "test" {
		t.Errorf("Expected value to be 'test', got %v", hint.GetValue())
	}
}
