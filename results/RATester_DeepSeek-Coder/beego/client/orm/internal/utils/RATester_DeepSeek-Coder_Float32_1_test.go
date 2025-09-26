package utils

import (
	"fmt"
	"testing"
)

func TestFloat32_1(t *testing.T) {
	tests := []struct {
		name    string
		s       StrTo
		want    float32
		wantErr bool
	}{
		{
			name:    "valid float32",
			s:       "3.14",
			want:    3.14,
			wantErr: false,
		},
		{
			name:    "invalid float32",
			s:       "invalid",
			want:    0,
			wantErr: true,
		},
		{
			name:    "out of range float32",
			s:       "3.14e+39",
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

			got, err := tt.s.Float32()
			if (err != nil) != tt.wantErr {
				t.Errorf("StrTo.Float32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StrTo.Float32() = %v, want %v", got, tt.want)
			}
		})
	}
}
