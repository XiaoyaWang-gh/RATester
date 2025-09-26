package cssjs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestdecodeTailwindCSSOptions_1(t *testing.T) {
	tests := []struct {
		name    string
		input   map[string]any
		want    TailwindCSSOptions
		wantErr bool
	}{
		{
			name:    "nil input",
			input:   nil,
			want:    TailwindCSSOptions{},
			wantErr: false,
		},
		{
			name: "valid input",
			input: map[string]any{
				"Minify":   true,
				"Optimize": true,
				"InlineImports": map[string]any{
					"SkipInlineImportsNotFound": true,
				},
			},
			want: TailwindCSSOptions{
				Minify:   true,
				Optimize: true,
				InlineImports: InlineImports{
					SkipInlineImportsNotFound: true,
				},
			},
			wantErr: false,
		},
		{
			name: "invalid input",
			input: map[string]any{
				"Minify":   "invalid",
				"Optimize": true,
				"InlineImports": map[string]any{
					"SkipInlineImportsNotFound": "invalid",
				},
			},
			want:    TailwindCSSOptions{},
			wantErr: true,
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeTailwindCSSOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
