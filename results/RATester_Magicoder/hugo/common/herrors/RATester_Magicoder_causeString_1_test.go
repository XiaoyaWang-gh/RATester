package herrors

import (
	"errors"
	"fmt"
	"testing"

	godartsassv1 "github.com/bep/godartsass"
	"github.com/bep/godartsass/v2"
	"github.com/bep/golibsass/libsass/libsasserrors"
)

func TestcauseString_1(t *testing.T) {
	tests := []struct {
		name  string
		cause error
		want  string
	}{
		{
			name: "godartsass.SassError",
			cause: godartsass.SassError{
				Message: "Test error message",
			},
			want: "Test error message",
		},
		{
			name: "godartsassv1.SassError",
			cause: godartsassv1.SassError{
				Message: "Test error message",
			},
			want: "Test error message",
		},
		{
			name: "libsasserrors.Error",
			cause: libsasserrors.Error{
				Message: "Test error message",
			},
			want: "Test error message",
		},
		{
			name:  "default",
			cause: errors.New("Test error message"),
			want:  "Test error message",
		},
		{
			name:  "nil",
			cause: nil,
			want:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			e := &fileError{
				cause: tt.cause,
			}
			if got := e.causeString(); got != tt.want {
				t.Errorf("causeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
