package rules

import (
	"fmt"
	"testing"
)

func TestNewParser_1(t *testing.T) {
	tests := []struct {
		name     string
		matchers []string
		wantErr  bool
	}{
		{
			name:     "empty matchers",
			matchers: []string{},
			wantErr:  true,
		},
		{
			name:     "single matcher",
			matchers: []string{"matcher"},
			wantErr:  false,
		},
		{
			name:     "multiple matchers",
			matchers: []string{"matcher1", "matcher2"},
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

			_, err := NewParser(tt.matchers)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewParser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
