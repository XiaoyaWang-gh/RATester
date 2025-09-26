package template

import (
	"fmt"
	"testing"
)

func TestErrorf_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &state{
		tmpl: &Template{
			name: "test",
		},
	}

	testFormat := "test format"
	testArgs := []any{"arg1", "arg2"}

	expectedFormat := fmt.Sprintf("template: %s: %s", doublePercent(s.tmpl.Name()), testFormat)
	expectedError := fmt.Errorf(expectedFormat, testArgs...)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but no panic occurred")
		} else {
			err, ok := r.(ExecError)
			if !ok {
				t.Errorf("Expected ExecError, but got %T", r)
			}
			if err.Name != s.tmpl.Name() {
				t.Errorf("Expected Name %s, but got %s", s.tmpl.Name(), err.Name)
			}
			if err.Err.Error() != expectedError.Error() {
				t.Errorf("Expected Error %s, but got %s", expectedError.Error(), err.Err.Error())
			}
		}
	}()

	s.errorf(testFormat, testArgs...)
}
