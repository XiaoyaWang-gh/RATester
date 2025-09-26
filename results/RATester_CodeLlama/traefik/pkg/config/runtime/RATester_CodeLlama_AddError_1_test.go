package runtime

import (
	"errors"
	"fmt"
	"testing"
)

func TestAddError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &UDPServiceInfo{}
	err := errors.New("test error")
	critical := true

	s.AddError(err, critical)

	if len(s.Err) != 1 {
		t.Errorf("expected 1 error, got %d", len(s.Err))
	}

	if s.Err[0] != err.Error() {
		t.Errorf("expected error %q, got %q", err.Error(), s.Err[0])
	}

	if s.Status != StatusDisabled {
		t.Errorf("expected status %q, got %q", StatusDisabled, s.Status)
	}
}
