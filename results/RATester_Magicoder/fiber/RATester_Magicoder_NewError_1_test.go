package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewError_1(t *testing.T) {
	tests := []struct {
		name    string
		code    int
		message string
		want    *Error
	}{
		{
			name:    "Test case 1",
			code:    400,
			message: "Bad Request",
			want:    &Error{Message: "Bad Request", Code: 400},
		},
		{
			name:    "Test case 2",
			code:    404,
			message: "Not Found",
			want:    &Error{Message: "Not Found", Code: 404},
		},
		{
			name:    "Test case 3",
			code:    500,
			message: "Internal Server Error",
			want:    &Error{Message: "Internal Server Error", Code: 500},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewError(tt.code, tt.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewError() = %v, want %v", got, tt.want)
			}
		})
	}
}
