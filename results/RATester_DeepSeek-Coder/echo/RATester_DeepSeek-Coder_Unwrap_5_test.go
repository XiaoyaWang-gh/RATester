package echo

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestUnwrap_5(t *testing.T) {
	tests := []struct {
		name     string
		he       *HTTPError
		expected error
	}{
		{
			name: "Test Unwrap with internal error",
			he: &HTTPError{
				Internal: errors.New("internal error"),
			},
			expected: errors.New("internal error"),
		},
		{
			name: "Test Unwrap without internal error",
			he: &HTTPError{
				Internal: nil,
			},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.he.Unwrap()
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("HTTPError.Unwrap() = %v, want %v", got, tt.expected)
			}
		})
	}
}
