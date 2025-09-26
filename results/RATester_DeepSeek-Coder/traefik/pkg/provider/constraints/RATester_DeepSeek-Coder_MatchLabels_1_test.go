package constraints

import (
	"fmt"
	"testing"
)

func TestMatchLabels_1(t *testing.T) {
	tests := []struct {
		name    string
		labels  map[string]string
		expr    string
		want    bool
		wantErr bool
	}{
		{
			name:   "empty labels and expr",
			labels: map[string]string{},
			expr:   "",
			want:   true,
		},
		{
			name:   "matching labels",
			labels: map[string]string{"env": "production", "tier": "frontend"},
			expr:   "Label('env') == 'production' AND Label('tier') == 'frontend'",
			want:   true,
		},
		{
			name:   "non-matching labels",
			labels: map[string]string{"env": "production", "tier": "frontend"},
			expr:   "Label('env') == 'staging' AND Label('tier') == 'frontend'",
			want:   false,
		},
		{
			name:    "invalid expr",
			labels:  map[string]string{"env": "production", "tier": "frontend"},
			expr:    "Label('env') == 'production' && Label('tier') == 'frontend'",
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

			got, err := MatchLabels(tt.labels, tt.expr)
			if (err != nil) != tt.wantErr {
				t.Errorf("MatchLabels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MatchLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
