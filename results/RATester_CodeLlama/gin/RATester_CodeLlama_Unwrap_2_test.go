package gin

import (
	"errors"
	"fmt"
	"testing"
)

func TestUnwrap_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	err := errors.New("test")
	msg := &Error{
		Err: err,
	}
	if msg.Unwrap() != err {
		t.Errorf("msg.Unwrap() = %v, want %v", msg.Unwrap(), err)
	}
}
