package echo

import (
	"errors"
	"fmt"
	"testing"
)

func TestError_2(t *testing.T) {
	tests := []struct {
		name string
		he   *HTTPError
		want string
	}{
		{
			name: "Test with nil internal error",
			he: &HTTPError{
				Code:     500,
				Message:  "Internal Server Error",
				Internal: nil,
			},
			want: "code=500, message=Internal Server Error",
		},
		{
			name: "Test with non-nil internal error",
			he: &HTTPError{
				Code:     500,
				Message:  "Internal Server Error",
				Internal: errors.New("External dependency error"),
			},
			want: "code=500, message=Internal Server Error, internal=External dependency error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.he.Error(); got != tt.want {
				t.Errorf("HTTPError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
