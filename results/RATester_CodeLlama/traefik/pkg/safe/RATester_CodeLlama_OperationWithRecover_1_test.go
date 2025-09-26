package safe

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestOperationWithRecover_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	operation := func() error {
		return errors.New("operation error")
	}
	operationWithRecover := OperationWithRecover(operation)
	err := operationWithRecover()
	if err == nil {
		t.Error("expected error")
	}
	if !strings.Contains(err.Error(), "operation error") {
		t.Errorf("expected error to contain 'operation error', got: %s", err)
	}
	if !strings.Contains(err.Error(), "panic in operation") {
		t.Errorf("expected error to contain 'panic in operation', got: %s", err)
	}
}
