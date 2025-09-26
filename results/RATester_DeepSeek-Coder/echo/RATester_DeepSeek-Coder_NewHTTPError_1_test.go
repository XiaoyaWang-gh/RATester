package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewHTTPError_1(t *testing.T) {
	t.Run("TestNewHTTPError", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		tests := []struct {
			name    string
			code    int
			message interface{}
			want    *HTTPError
		}{
			{
				name:    "Test with 404 status code",
				code:    404,
				message: "Not Found",
				want:    &HTTPError{Code: 404, Message: "Not Found"},
			},
			{
				name:    "Test with 500 status code",
				code:    500,
				message: "Internal Server Error",
				want:    &HTTPError{Code: 500, Message: "Internal Server Error"},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				if got := NewHTTPError(tt.code, tt.message); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("NewHTTPError() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}
