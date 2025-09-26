package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewHTTPError_1(t *testing.T) {
	tests := []struct {
		name    string
		code    int
		message string
		want    *HTTPError
	}{
		{
			name:    "Test case 1",
			code:    400,
			message: "Bad Request",
			want:    &HTTPError{Code: 400, Message: "Bad Request"},
		},
		{
			name:    "Test case 2",
			code:    404,
			message: "Not Found",
			want:    &HTTPError{Code: 404, Message: "Not Found"},
		},
		{
			name:    "Test case 3",
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
}
