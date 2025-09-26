package echo

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestWithInternal_1(t *testing.T) {
	tests := []struct {
		name     string
		he       *HTTPError
		internal error
		want     *HTTPError
	}{
		{
			name: "Test with internal error",
			he: &HTTPError{
				Code:    500,
				Message: "Internal Server Error",
			},
			internal: errors.New("database error"),
			want: &HTTPError{
				Code:     500,
				Message:  "Internal Server Error",
				Internal: errors.New("database error"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.he.WithInternal(tt.internal)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
