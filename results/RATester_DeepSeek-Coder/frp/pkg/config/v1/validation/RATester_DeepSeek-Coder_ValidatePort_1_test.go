package validation

import (
	"fmt"
	"strings"
	"testing"
)

func TestValidatePort_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		port       int
		fieldPath  string
		expectErr  bool
		errorMatch string
	}

	tests := []test{
		{
			port:       65536,
			fieldPath:  "port",
			expectErr:  true,
			errorMatch: "port number 65536 must be in the range 0..65535",
		},
		{
			port:       -1,
			fieldPath:  "port",
			expectErr:  true,
			errorMatch: "port number -1 must be in the range 0..65535",
		},
		{
			port:       0,
			fieldPath:  "port",
			expectErr:  false,
			errorMatch: "",
		},
		{
			port:       65535,
			fieldPath:  "port",
			expectErr:  false,
			errorMatch: "",
		},
		{
			port:       30000,
			fieldPath:  "port",
			expectErr:  false,
			errorMatch: "",
		},
	}

	for _, test := range tests {
		err := ValidatePort(test.port, test.fieldPath)
		if test.expectErr && err == nil {
			t.Errorf("Expected error for port %d, fieldPath %s, but got nil", test.port, test.fieldPath)
		} else if !test.expectErr && err != nil {
			t.Errorf("Expected no error for port %d, fieldPath %s, but got error: %v", test.port, test.fieldPath, err)
		}

		if err != nil && !strings.Contains(err.Error(), test.errorMatch) {
			t.Errorf("Expected error message to contain '%s', but got '%s'", test.errorMatch, err.Error())
		}
	}
}
