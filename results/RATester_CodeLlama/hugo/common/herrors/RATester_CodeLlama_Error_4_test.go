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

	e := TextSegmentError{
		Segment: "foo",
		Err:     errors.New("bar"),
	}
	if e.Error() != "foo: bar" {
		t.Errorf("expected %q, got %q", "foo: bar", e.Error())
	}
}
