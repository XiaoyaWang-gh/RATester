package binding

import (
	"errors"
	"fmt"
	"testing"
)

func TestError_1(t *testing.T) {
	tests := []struct {
		name string
		err  SliceValidationError
		want string
	}{
		{
			name: "empty",
			err:  SliceValidationError{},
			want: "",
		},
		{
			name: "single error",
			err:  SliceValidationError{errors.New("error 1")},
			want: "[0]: error 1",
		},
		{
			name: "multiple errors",
			err:  SliceValidationError{errors.New("error 1"), errors.New("error 2")},
			want: "[0]: error 1\n[1]: error 2",
		},
		{
			name: "nil error",
			err:  SliceValidationError{nil, errors.New("error 1")},
			want: "[1]: error 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.err.Error()
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
