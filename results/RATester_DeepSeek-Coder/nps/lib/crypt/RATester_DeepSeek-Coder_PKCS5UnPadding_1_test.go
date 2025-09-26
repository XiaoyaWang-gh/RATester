package crypt

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestPKCS5UnPadding_1(t *testing.T) {
	t.Run("TestPKCS5UnPadding", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testCases := []struct {
			name     string
			input    []byte
			expected []byte
			err      error
		}{
			{
				name:     "Test case 1",
				input:    []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x01},
				expected: []byte{0x01, 0x02, 0x03, 0x04, 0x05},
				err:      nil,
			},
			{
				name:     "Test case 2",
				input:    []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x00},
				expected: []byte{0x01, 0x02, 0x03, 0x04, 0x05},
				err:      nil,
			},
			{
				name:     "Test case 3",
				input:    []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x05},
				expected: nil,
				err:      errors.New("len error"),
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				err, actual := PKCS5UnPadding(tc.input)
				if !reflect.DeepEqual(actual, tc.expected) {
					t.Errorf("Expected %v, got %v", tc.expected, actual)
				}
				if !errors.Is(err, tc.err) {
					t.Errorf("Expected error %v, got %v", tc.err, err)
				}
			})
		}
	})
}
