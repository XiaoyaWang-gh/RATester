package echo

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestNewBindingError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	sourceParam := "testParam"
	values := []string{"value1", "value2"}
	message := "test message"
	internalError := errors.New("internal error")

	err := NewBindingError(sourceParam, values, message, internalError)

	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	bindingError, ok := err.(*BindingError)
	if !ok {
		t.Errorf("Expected error of type *BindingError, got %T", err)
		return
	}

	if bindingError.Field != sourceParam {
		t.Errorf("Expected Field to be %s, got %s", sourceParam, bindingError.Field)
	}

	if !reflect.DeepEqual(bindingError.Values, values) {
		t.Errorf("Expected Values to be %v, got %v", values, bindingError.Values)
	}

	if bindingError.HTTPError.Message != message {
		t.Errorf("Expected HTTPError.Message to be %s, got %s", message, bindingError.HTTPError.Message)
	}

	if bindingError.HTTPError.Internal.Error() != internalError.Error() {
		t.Errorf("Expected HTTPError.Internal to be %s, got %s", internalError.Error(), bindingError.HTTPError.Internal.Error())
	}
}
