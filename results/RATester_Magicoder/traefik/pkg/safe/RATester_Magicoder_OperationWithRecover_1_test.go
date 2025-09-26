package safe

import (
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
		panic("test panic")
	}

	wrappedOperation := OperationWithRecover(operation)
	err := wrappedOperation()

	if err == nil {
		t.Error("Expected error, got nil")
	}

	if !strings.Contains(err.Error(), "panic in operation") {
		t.Errorf("Expected error to contain 'panic in operation', got '%s'", err.Error())
	}
}
