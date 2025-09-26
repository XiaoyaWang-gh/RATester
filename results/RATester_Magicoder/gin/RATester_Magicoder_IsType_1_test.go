package gin

import (
	"fmt"
	"testing"
)

func TestIsType_1(t *testing.T) {
	testCases := []struct {
		name  string
		msg   *Error
		flags ErrorType
		want  bool
	}{
		{
			name: "Test case 1",
			msg: &Error{
				Type: ErrorType(1),
			},
			flags: ErrorType(1),
			want:  true,
		},
		{
			name: "Test case 2",
			msg: &Error{
				Type: ErrorType(2),
			},
			flags: ErrorType(1),
			want:  false,
		},
		{
			name: "Test case 3",
			msg: &Error{
				Type: ErrorType(3),
			},
			flags: ErrorType(3),
			want:  true,
		},
		{
			name: "Test case 4",
			msg: &Error{
				Type: ErrorType(4),
			},
			flags: ErrorType(5),
			want:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tc.msg.IsType(tc.flags)
			if got != tc.want {
				t.Errorf("got %t, want %t", got, tc.want)
			}
		})
	}
}
