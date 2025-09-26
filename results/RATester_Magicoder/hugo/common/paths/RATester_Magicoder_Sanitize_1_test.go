package paths

import (
	"fmt"
	"testing"
)

func TestSanitize_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "Test case 1",
			arg:  "test-case-1",
			want: "test-case-1",
		},
		{
			name: "Test case 2",
			arg:  "test case 2",
			want: "test-case-2",
		},
		{
			name: "Test case 3",
			arg:  "test-case-3",
			want: "test-case-3",
		},
		{
			name: "Test case 4",
			arg:  "test case 4",
			want: "test-case-4",
		},
		{
			name: "Test case 5",
			arg:  "test-case-5",
			want: "test-case-5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Sanitize(tt.arg); got != tt.want {
				t.Errorf("Sanitize() = %v, want %v", got, tt.want)
			}
		})
	}
}
