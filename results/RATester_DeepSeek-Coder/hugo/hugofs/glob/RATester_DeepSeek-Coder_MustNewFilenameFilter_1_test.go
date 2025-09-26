package glob

import (
	"fmt"
	"testing"
)

func TestMustNewFilenameFilter_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		inclusions []string
		exclusions []string
		wantErr    bool
		wantPanic  bool
	}{
		{
			name:       "valid inclusions and exclusions",
			inclusions: []string{"*.txt", "*.go"},
			wantErr:    false,
			wantPanic:  false,
		},
		{
			name:       "invalid inclusions",
			inclusions: []string{"[*.txt", "*.go"},
			wantErr:    true,
			wantPanic:  true,
		},
		{
			name:       "invalid exclusions",
			inclusions: []string{"*.txt", "*.go"},
			exclusions: []string{"[*.txt", "*.go"},
			wantErr:    true,
			wantPanic:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			defer func() {
				if r := recover(); (r != nil) != tt.wantPanic {
					t.Errorf("MustNewFilenameFilter() = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()

			_, err := NewFilenameFilter(tt.inclusions, tt.exclusions)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFilenameFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
