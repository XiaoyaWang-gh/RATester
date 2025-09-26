package page

import (
	"errors"
	"fmt"
	"testing"
)

func TestError_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pee := &permalinkExpandError{
		pattern: "test_pattern",
		err:     errors.New("test_error"),
	}

	expected := fmt.Sprintf("error expanding %q: %s", pee.pattern, pee.err)
	result := pee.Error()

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}
