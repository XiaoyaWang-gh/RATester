package echo

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewBindingError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	sourceParam := "sourceParam"
	values := []string{"value1", "value2"}
	message := "message"
	internalError := errors.New("internalError")
	err := NewBindingError(sourceParam, values, message, internalError)
	if err == nil {
		t.Errorf("NewBindingError() = %v, want %v", err, "NewBindingError() = %v, want %v")
	}
}
