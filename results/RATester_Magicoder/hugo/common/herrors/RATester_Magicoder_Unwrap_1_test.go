package herrors

import (
	"errors"
	"fmt"
	"testing"
)

func TestUnwrap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testErr := errors.New("test error")
	segmentErr := TextSegmentError{
		Segment: "test segment",
		Err:     testErr,
	}

	unwrappedErr := segmentErr.Unwrap()

	if unwrappedErr != testErr {
		t.Errorf("Expected unwrapped error to be %v, but got %v", testErr, unwrappedErr)
	}
}
