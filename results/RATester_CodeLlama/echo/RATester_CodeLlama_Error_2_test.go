package echo

import (
	"fmt"
	"testing"
)

func TestError_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	he := &HTTPError{
		Code:    404,
		Message: "not found",
	}
	if he.Error() != "code=404, message=not found" {
		t.Errorf("expected %q, got %q", "code=404, message=not found", he.Error())
	}
}
