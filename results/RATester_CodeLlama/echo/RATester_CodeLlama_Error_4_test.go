package echo

import (
	"fmt"
	"testing"
)

func TestError_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	be := &BindingError{
		Field: "field",
		HTTPError: &HTTPError{
			Code:    400,
			Message: "message",
		},
	}
	if be.Error() != "message, field=field" {
		t.Errorf("Error() = %v, want %v", be.Error(), "message, field=field")
	}
}
