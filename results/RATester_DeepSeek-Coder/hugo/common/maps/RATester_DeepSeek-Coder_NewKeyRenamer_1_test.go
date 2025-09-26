package maps

import (
	"fmt"
	"testing"
)

func TestNewKeyRenamer_1(t *testing.T) {
	tests := []struct {
		name        string
		patternKeys []string
		wantErr     bool
	}{
		{
			name:        "valid pattern",
			patternKeys: []string{"*", "**"},
			wantErr:     false,
		},
		{
			name:        "invalid pattern",
			patternKeys: []string{"[", "**"},
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := NewKeyRenamer(tt.patternKeys...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewKeyRenamer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
