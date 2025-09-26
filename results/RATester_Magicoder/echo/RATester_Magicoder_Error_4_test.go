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
		Field: "testField",
		HTTPError: &HTTPError{
			Message: "testMessage",
			Code:    400,
		},
	}

	expected := "testMessage, field=testField"
	actual := be.Error()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
