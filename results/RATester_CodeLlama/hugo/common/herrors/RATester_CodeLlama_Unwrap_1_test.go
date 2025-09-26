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

	e := TextSegmentError{
		Segment: "foo",
		Err:     errors.New("bar"),
	}
	if got := e.Unwrap(); got != e.Err {
		t.Errorf("TextSegmentError.Unwrap() = %v, want %v", got, e.Err)
	}
}
