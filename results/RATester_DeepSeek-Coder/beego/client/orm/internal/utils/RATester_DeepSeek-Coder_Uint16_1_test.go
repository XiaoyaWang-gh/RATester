package utils

import (
	"fmt"
	"testing"
)

func TestUint16_1(t *testing.T) {
	tests := []struct {
		name    string
		s       StrTo
		want    uint16
		wantErr bool
	}{
		{
			name:    "valid",
			s:       "12345",
			want:    12345,
			wantErr: false,
		},
		{
			name:    "invalid",
			s:       "abc",
			want:    0,
			wantErr: true,
		},
		{
			name:    "out of range",
			s:       "65536",
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

			got, err := tt.s.Uint16()
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint16() got = %v, want %v", got, tt.want)
			}
		})
	}
}
