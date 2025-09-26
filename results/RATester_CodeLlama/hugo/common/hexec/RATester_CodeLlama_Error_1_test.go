package hexec

import (
	"fmt"
	"testing"
)

func TestError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &NotFoundError{
		name:   "foo",
		method: "bar",
	}
	if e.Error() != "binary with name \"foo\" not found bar" {
		t.Errorf("expected \"binary with name \"foo\" not found bar\", got %q", e.Error())
	}
}
