package utils

import (
	"fmt"
	"testing"
)

func TestFloat64_2(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		s       StrTo
		want    float64
		wantErr bool
	}{
		{
			name:    "valid float",
			s:       "3.14159",
			want:    3.14159,
			wantErr: false,
		},
		{
			name:    "invalid float",
			s:       "invalid",
			want:    0,
			wantErr: true,
		},
		{
			name:    "empty string",
			s:       "",
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

			got, err := tt.s.Float64()
			if (err != nil) != tt.wantErr {
				t.Errorf("StrTo.Float64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StrTo.Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}
