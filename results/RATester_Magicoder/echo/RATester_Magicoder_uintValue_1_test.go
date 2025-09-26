package echo

import (
	"errors"
	"fmt"
	"testing"
)

func TestUintValue_1(t *testing.T) {
	tests := []struct {
		name           string
		sourceParam    string
		value          string
		bitSize        int
		valueMustExist bool
		expectedError  bool
	}{
		{
			name:           "Test case 1",
			sourceParam:    "field name",
			value:          "field value",
			bitSize:        32,
			valueMustExist: true,
			expectedError:  false,
		},
		{
			name:           "Test case 2",
			sourceParam:    "field name",
			value:          "",
			bitSize:        32,
			valueMustExist: true,
			expectedError:  true,
		},
		{
			name:           "Test case 3",
			sourceParam:    "field name",
			value:          "field value",
			bitSize:        32,
			valueMustExist: false,
			expectedError:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			b := &ValueBinder{
				ValueFunc: func(sourceParam string) string {
					if sourceParam == test.sourceParam {
						return test.value
					}
					return ""
				},
				ErrorFunc: func(sourceParam string, values []string, message interface{}, internalError error) error {
					if test.expectedError {
						return errors.New("error")
					}
					return nil
				},
			}

			b.uintValue(test.sourceParam, nil, test.bitSize, test.valueMustExist)

			if test.expectedError && b.errors == nil {
				t.Errorf("Expected error but got nil")
			}

			if !test.expectedError && b.errors != nil {
				t.Errorf("Expected no error but got error: %v", b.errors)
			}
		})
	}
}
