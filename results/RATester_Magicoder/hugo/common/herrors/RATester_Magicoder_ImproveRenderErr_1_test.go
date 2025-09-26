package herrors

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestImproveRenderErr_1(t *testing.T) {
	tests := []struct {
		name     string
		inErr    error
		expected error
	}{
		{
			name:     "nil error",
			inErr:    nil,
			expected: nil,
		},
		{
			name:     "non-nil error",
			inErr:    errors.New("some error"),
			expected: &errMessage{msg: "some error", err: nil},
		},
		{
			name:     "deferred error",
			inErr:    errors.New("deferred: some error"),
			expected: &errMessage{msg: "executing some error", err: nil},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			outErr := ImproveRenderErr(test.inErr)
			if !reflect.DeepEqual(outErr, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, outErr)
			}
		})
	}
}
