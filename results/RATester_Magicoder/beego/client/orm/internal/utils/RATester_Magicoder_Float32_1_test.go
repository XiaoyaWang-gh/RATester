package utils

import (
	"fmt"
	"testing"
)

func TestFloat32_1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    float32
		wantErr bool
	}{
		{
			name:    "valid float32",
			input:   "3.14",
			want:    3.14,
			wantErr: false,
		},
		{
			name:    "invalid float32",
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

			f := StrTo(tt.input)
			got, err := f.Float32()
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
