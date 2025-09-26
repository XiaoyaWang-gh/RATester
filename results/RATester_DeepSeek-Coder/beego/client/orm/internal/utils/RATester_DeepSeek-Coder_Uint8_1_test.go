package utils

import (
	"fmt"
	"testing"
)

func TestUint8_1(t *testing.T) {
	tests := []struct {
		name    string
		s       StrTo
		want    uint8
		wantErr bool
	}{
		{
			name:    "valid uint8",
			s:       "255",
			want:    255,
			wantErr: false,
		},
		{
			name:    "invalid uint8",
			s:       "256",
			want:    0,
			wantErr: true,
		},
		{
			name:    "empty string",
			s:       "",
			want:    0,
			wantErr: true,
		},
		{
			name:    "non-numeric string",
			s:       "abc",
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

			got, err := tt.s.Uint8()
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint8() got = %v, want %v", got, tt.want)
			}
		})
	}
}
