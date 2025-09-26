package schema

import (
	"errors"
	"fmt"
	"testing"
)

func TestError_3(t *testing.T) {
	tests := []struct {
		name string
		e    MultiError
		want string
	}{
		{
			name: "no errors",
			e:    MultiError{},
			want: "(0 errors)",
		},
		{
			name: "one error",
			e:    MultiError{"field": errors.New("error 1")},
			want: "error 1",
		},
		{
			name: "two errors",
			e:    MultiError{"field1": errors.New("error 1"), "field2": errors.New("error 2")},
			want: "error 1 (and 1 other error)",
		},
		{
			name: "three errors",
			e:    MultiError{"field1": errors.New("error 1"), "field2": errors.New("error 2"), "field3": errors.New("error 3")},
			want: "error 1 (and 2 other errors)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.e.Error(); got != tt.want {
				t.Errorf("MultiError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
