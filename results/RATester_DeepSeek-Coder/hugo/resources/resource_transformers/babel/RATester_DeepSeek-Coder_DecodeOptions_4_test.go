package babel

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDecodeOptions_4(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   map[string]any
		want    Options
		wantErr bool
	}{
		{
			name:  "nil input",
			input: nil,
			want:  Options{},
		},
		{
			name:  "empty input",
			input: map[string]any{},
			want:  Options{},
		},
		{
			name: "minified",
			input: map[string]any{
				"Minified": true,
			},
			want: Options{
				Minified: true,
			},
		},
		{
			name: "verbose",
			input: map[string]any{
				"Verbose": true,
			},
			want: Options{
				Verbose: true,
			},
		},
		{
			name: "complex",
			input: map[string]any{
				"Config":   "/path/to/config",
				"Minified": true,
				"Verbose":  true,
			},
			want: Options{
				Config:   "/path/to/config",
				Minified: true,
				Verbose:  true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := DecodeOptions(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeOptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
