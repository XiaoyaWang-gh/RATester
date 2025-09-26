package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConvertFloat64_1(t *testing.T) {
	t.Run("Test convertFloat64 function", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testCases := []struct {
			name  string
			value string
			want  reflect.Value
		}{
			{
				name:  "Test with valid float string",
				value: "3.14",
				want:  reflect.ValueOf(3.14),
			},
			{
				name:  "Test with invalid float string",
				value: "invalid",
				want:  reflect.ValueOf(0),
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				got := convertFloat64(tc.value)
				if got.Float() != tc.want.Float() {
					t.Errorf("Expected %v, got %v", tc.want, got)
				}
			})
		}
	})
}
