package echo

import (
	"errors"
	"fmt"
	"testing"
)

func TestMustCustomFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &ValueBinder{}
	sourceParam := "testParam"
	customFunc := func(values []string) []error {
		return []error{errors.New("test error")}
	}

	b.MustCustomFunc(sourceParam, customFunc)

	if len(b.errors) != 1 {
		t.Errorf("Expected 1 error, got %d", len(b.errors))
	}
}
