package herrors

import (
	"fmt"
	"testing"
)

func TestErrorContext_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &fileError{}
	if e.ErrorContext() != nil {
		t.Errorf("ErrorContext() should be nil")
	}
}
