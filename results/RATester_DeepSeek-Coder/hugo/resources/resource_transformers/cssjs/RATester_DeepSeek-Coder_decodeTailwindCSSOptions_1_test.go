package cssjs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDecodeTailwindCSSOptions_1(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		input    map[string]any
		expected TailwindCSSOptions
		wantErr  bool
	}

	tests := []testCase{
		{
			name: "minify",
			input: map[string]any{
				"minify": true,
			},
			expected: TailwindCSSOptions{
				Minify: true,
			},
			wantErr: false,
		},
		{
			name: "optimize",
			input: map[string]any{
				"optimize": true,
			},
			expected: TailwindCSSOptions{
				Optimize: true,
			},
			wantErr: false,
		},
		{
			name: "invalid key",
			input: map[string]any{
				"invalid": true,
			},
			expected: TailwindCSSOptions{},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := decodeTailwindCSSOptions(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeTailwindCSSOptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("decodeTailwindCSSOptions() = %v, want %v", got, tt.expected)
			}
		})
	}
}
