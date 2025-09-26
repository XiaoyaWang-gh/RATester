package math

import (
	"fmt"
	"testing"
)

func TestAbs_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    float64
		wantErr bool
	}{
		{
			name:    "Test with positive number",
			input:   5,
			want:    5,
			wantErr: false,
		},
		{
			name:    "Test with negative number",
			input:   -3,
			want:    3,
			wantErr: false,
		},
		{
			name:    "Test with zero",
			input:   0,
			want:    0,
			wantErr: false,
		},
		{
			name:    "Test with non-numeric input",
			input:   "string",
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

			got, err := ns.Abs(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Abs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
