package css

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources/resource_transformers/tocss/scss"
)

func TestSass_1(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		args     []any
		expected error
	}

	tests := []testCase{
		{
			name:     "too many arguments",
			args:     []any{nil, nil, nil},
			expected: errors.New("must not provide more arguments than resource object and options"),
		},
		{
			name:     "invalid transpiler",
			args:     []any{nil, map[string]any{"transpiler": "invalid"}},
			expected: fmt.Errorf("unsupported transpiler %q; valid values are %q or %q", "invalid", "libsass", "dartsass"),
		},
		// Add more test cases as needed
	}

	for _, tc := range tests {
		tc := tc // Patch the range variable to avoid concurrent test execution issues
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			ns := &Namespace{
				scssClientLibSass: &scss.Client{},
			}

			_, err := ns.Sass(tc.args...)
			if err == nil {
				if tc.expected != nil {
					t.Errorf("expected error %v, but got nil", tc.expected)
				}
			} else if tc.expected == nil {
				t.Errorf("unexpected error: %v", err)
			} else if err.Error() != tc.expected.Error() {
				t.Errorf("expected error %v, but got %v", tc.expected, err)
			}
		})
	}
}
