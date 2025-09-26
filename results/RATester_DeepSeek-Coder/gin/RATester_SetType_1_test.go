package gin

import (
	"errors"
	"fmt"
	"testing"
)

func TestSetType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testErr := &Error{
		Err:  errors.New("test error"),
		Type: 0,
		Meta: "test meta",
	}

	newErr := testErr.SetType(1)

	if newErr.Type != 1 {
		t.Errorf("Expected Type to be 1, got %v", newErr.Type)
	}
}
