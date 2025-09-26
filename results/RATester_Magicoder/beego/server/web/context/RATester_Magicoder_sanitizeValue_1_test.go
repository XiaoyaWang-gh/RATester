package context

import (
	"fmt"
	"testing"
)

func TestsanitizeValue_1(t *testing.T) {
	tests := []struct {
		name string
		v    string
		want string
	}{
		{
			name: "Test case 1",
			v:    "<script>alert('XSS')</script>",
			want: "alert('XSS')",
		},
		{
			name: "Test case 2",
			v:    "<img src=x onerror=alert('XSS')>",
			want: "<img src=x>",
		},
		{
			name: "Test case 3",
			v:    "<iframe src=javascript:alert('XSS')>",
			want: "<iframe>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := sanitizeValue(tt.v); got != tt.want {
				t.Errorf("sanitizeValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
