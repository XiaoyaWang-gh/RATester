package validation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetDefaultMessage_1(t *testing.T) {
	t.Run("TestSetDefaultMessage", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testCases := []struct {
			name  string
			input map[string]string
			want  map[string]string
		}{
			{
				name:  "Test case 1: Normal case",
				input: map[string]string{"key1": "value1", "key2": "value2"},
				want:  map[string]string{"key1": "value1", "key2": "value2"},
			},
			{
				name:  "Test case 2: Empty map",
				input: map[string]string{},
				want:  map[string]string{},
			},
			{
				name:  "Test case 3: Nil map",
				input: nil,
				want:  map[string]string{},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				SetDefaultMessage(tc.input)
				if !reflect.DeepEqual(MessageTmpls, tc.want) {
					t.Errorf("Expected %v, got %v", tc.want, MessageTmpls)
				}
			})
		}
	})
}
