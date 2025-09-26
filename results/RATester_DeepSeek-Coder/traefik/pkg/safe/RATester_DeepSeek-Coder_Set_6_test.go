package safe

import (
	"fmt"
	"testing"
)

func TestSet_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Safe{}
	testValue := "test"
	s.Set(testValue)

	if s.value != testValue {
		t.Errorf("Expected %v, got %v", testValue, s.value)
	}
}
