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
		pattern: "foo",
		err:     errors.New("bar"),
	}
	if got := pee.Error(); got != "error expanding \"foo\": bar" {
		t.Errorf("pee.Error() = %q; want %q", got, "error expanding \"foo\": bar")
	}
}
