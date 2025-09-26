package fiber

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

	e := &Error{
		Message: "test",
	}
	if e.Error() != "test" {
		t.Errorf("Error() = %v, want %v", e.Error(), "test")
	}
}
