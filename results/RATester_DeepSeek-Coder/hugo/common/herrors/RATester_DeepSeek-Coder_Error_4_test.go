package herrors

import (
	"errors"
	"fmt"
	"testing"
)

func TestError_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testErr := TextSegmentError{
		Segment: "testSegment",
		Err:     errors.New("test error"),
	}

	expected := "testSegment: test error"
	actual := testErr.Error()

	if actual != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}
