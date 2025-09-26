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
			e:    MultiError{"field": errors.New("error message")},
			want: "error message",
		},
		{
			name: "two errors",
			e:    MultiError{"field1": errors.New("error message1"), "field2": errors.New("error message2")},
			want: "error message1 (and 1 other error)",
		},
		{
			name: "three errors",
			e:    MultiError{"field1": errors.New("error message1"), "field2": errors.New("error message2"), "field3": errors.New("error message3")},
			want: "error message1 (and 2 other errors)",
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
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
