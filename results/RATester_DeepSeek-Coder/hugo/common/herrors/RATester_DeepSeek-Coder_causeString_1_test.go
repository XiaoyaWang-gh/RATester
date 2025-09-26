package herrors

import (
	"errors"
	"fmt"
	"testing"

	godartsassv1 "github.com/bep/godartsass"
	"github.com/bep/godartsass/v2"
	"github.com/bep/golibsass/libsass/libsasserrors"
)

func TestCauseString_1(t *testing.T) {
	type testCase struct {
		name      string
		fileError *fileError
		want      string
	}

	tests := []testCase{
		{
			name: "Test case 1",
			fileError: &fileError{
				cause: errors.New("test error"),
			},
			want: "test error",
		},
		{
			name: "Test case 2",
			fileError: &fileError{
				cause: godartsass.SassError{Message: "test error"},
			},
			want: "test error",
		},
		{
			name: "Test case 3",
			fileError: &fileError{
				cause: godartsassv1.SassError{Message: "test error"},
			},
			want: "test error",
		},
		{
			name: "Test case 4",
			fileError: &fileError{
				cause: libsasserrors.Error{Message: "test error"},
			},
			want: "test error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.fileError.causeString(); got != tt.want {
				t.Errorf("causeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
