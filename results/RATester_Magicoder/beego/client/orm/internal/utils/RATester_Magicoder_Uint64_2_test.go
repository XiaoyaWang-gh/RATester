package utils

import (
	"fmt"
	"testing"
)

func TestUint64_2(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    uint64
		wantErr bool
	}{
		{
			name:    "valid uint64",
			input:   "1234567890",
			want:    1234567890,
			wantErr: false,
		},
		{
			name:    "invalid uint64",
			input:   "abc",
			want:    0,
			wantErr: true,
		},
		{
			name:    "out of range uint64",
			input:   "18446744073709551616",
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
			got, err := f.Uint64()
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint64() = %v, want %v", got, tt.want)
			}
		})
	}
}
