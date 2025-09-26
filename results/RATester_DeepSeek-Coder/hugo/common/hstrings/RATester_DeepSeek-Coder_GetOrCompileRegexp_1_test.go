package hstrings

import (
	"fmt"
	"testing"
)

func TestGetOrCompileRegexp_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		pattern string
		wantErr bool
	}{
		{
			name:    "valid regexp",
			pattern: "^[a-zA-Z0-9]*$",
			wantErr: false,
		},
		{
			name:    "invalid regexp",
			pattern: "*",
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

			_, err := GetOrCompileRegexp(tt.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrCompileRegexp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
