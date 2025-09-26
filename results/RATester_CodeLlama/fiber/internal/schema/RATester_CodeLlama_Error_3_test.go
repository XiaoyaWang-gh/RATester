package schema

import (
	"errors"
	"fmt"
	"testing"
)

func TestError_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := MultiError{
		"a": errors.New("a"),
		"b": errors.New("b"),
		"c": errors.New("c"),
	}
	if e.Error() != "a (and 2 other errors)" {
		t.Errorf("got %q, want %q", e.Error(), "a (and 2 other errors)")
	}
}
