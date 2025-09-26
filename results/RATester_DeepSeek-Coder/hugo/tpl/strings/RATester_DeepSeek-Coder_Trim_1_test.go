package strings

import (
	"fmt"
	"testing"
)

func TestTrim_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name     string
		s        any
		cutset   any
		expected string
		wantErr  bool
	}{
		{
			name:     "Trim string with string cutset",
			s:        "!!!Hello, Gophers!!!",
			cutset:   "!!!",
			expected: "Hello, Gophers",
			wantErr:  false,
		},
		{
			name:     "Trim string with non-string cutset",
			s:        "!!!Hello, Gophers!!!",
			cutset:   123,
			expected: "",
			wantErr:  true,
		},
		{
			name:     "Trim string with non-string s",
			s:        123,
			cutset:   "!!!",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "Trim string with empty cutset",
			s:        "Hello, Gophers!!!",
			cutset:   "",
			expected: "Hello, Gophers!!!",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ns.Trim(tt.s, tt.cutset)
			if (err != nil) != tt.wantErr {
				t.Errorf("Namespace.Trim() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("Namespace.Trim() = %v, want %v", got, tt.expected)
			}
		})
	}
}
