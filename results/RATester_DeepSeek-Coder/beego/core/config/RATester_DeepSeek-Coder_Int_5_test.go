package config

import (
	"errors"
	"fmt"
	"testing"
)

func TestInt_5(t *testing.T) {
	type testCase struct {
		name     string
		key      string
		expected int
		err      error
	}

	testCases := []testCase{
		{
			name:     "TestInt_Success",
			key:      "key1",
			expected: 123,
			err:      nil,
		},
		{
			name:     "TestInt_Error",
			key:      "key2",
			expected: 0,
			err:      errors.New("strconv.Atoi: parsing \"not_an_int\": invalid syntax"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			fc := &fakeConfigContainer{
				data: map[string]string{
					"key1": "123",
					"key2": "not_an_int",
				},
			}

			actual, err := fc.Int(tc.key)
			if tc.err != nil {
				if err == nil {
					t.Errorf("Expected error '%v', but got nil", tc.err)
				} else if err.Error() != tc.err.Error() {
					t.Errorf("Expected error '%v', but got '%v'", tc.err, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, but got '%v'", err)
				}
				if actual != tc.expected {
					t.Errorf("Expected '%v', but got '%v'", tc.expected, actual)
				}
			}
		})
	}
}
