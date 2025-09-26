package validation

import (
	"fmt"
	"testing"
)

func TestString_5(t *testing.T) {
	tests := []struct {
		name string
		e    *Error
		want string
	}{
		{
			name: "Test case 1",
			e: &Error{
				Message: "Test message",
			},
			want: "Test message",
		},
		{
			name: "Test case 2",
			e: &Error{
				Message: "Another test message",
			},
			want: "Another test message",
		},
		{
			name: "Test case 3",
			e: &Error{
				Message: "Yet another test message",
			},
			want: "Yet another test message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.e.String(); got != tt.want {
				t.Errorf("Error.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
