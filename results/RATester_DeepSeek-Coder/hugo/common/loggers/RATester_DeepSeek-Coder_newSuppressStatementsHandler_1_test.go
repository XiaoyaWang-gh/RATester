package loggers

import (
	"fmt"
	"testing"
)

func TestNewSuppressStatementsHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testStatements := map[string]bool{
		"statement1": true,
		"statement2": false,
	}

	handler := newSuppressStatementsHandler(testStatements)

	if handler == nil {
		t.Errorf("Expected handler to be created, but got nil")
	}

	if len(handler.statements) != len(testStatements) {
		t.Errorf("Expected handler to have %d statements, but got %d", len(testStatements), len(handler.statements))
	}

	for k, v := range testStatements {
		if handler.statements[k] != v {
			t.Errorf("Expected statement %s to be %t, but got %t", k, v, handler.statements[k])
		}
	}
}
