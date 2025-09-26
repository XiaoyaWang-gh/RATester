package math

import (
	"fmt"
	"testing"
)

func TestAbs_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    float64
		wantErr bool
	}{
		{
			name:    "Positive number",
			input:   5,
			want:    5,
			wantErr: false,
		},
		{
			name:    "Negative number",
			input:   -3,
			want:    3,
			wantErr: false,
		},
		{
			name:    "Zero",
			input:   0,
			want:    0,
			wantErr: false,
		},
		{
			name:    "Non-numeric input",
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
