package parse

import (
	"fmt"
	"strings"
	"testing"
)

func Testerrorf_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{}
	format := "test format"
	args := []any{"arg1", "arg2"}

	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if !ok {
				t.Errorf("Expected error, got %T", r)
			}
			if tree.Root != nil {
				t.Errorf("Expected Root to be nil, got %v", tree.Root)
			}
			if !strings.Contains(err.Error(), format) {
				t.Errorf("Expected error message to contain format, got %s", err.Error())
			}
		}
	}()

	tree.errorf(format, args...)
}
