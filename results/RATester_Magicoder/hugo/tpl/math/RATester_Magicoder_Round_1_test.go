package math

import (
	"fmt"
	"testing"
)

func TestRound_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    float64
		wantErr bool
	}{
		{
			name:    "Round with float64",
			input:   3.14,
			want:    3.0,
			wantErr: false,
		},
		{
			name:    "Round with int",
			input:   3,
			want:    3.0,
			wantErr: false,
		},
		{
			name:    "Round with string",
			input:   "3.14",
			want:    3.0,
			wantErr: false,
		},
		{
			name:    "Round with non-float value",
			input:   "abc",
			want:    0,
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

			got, err := ns.Round(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Round() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}
