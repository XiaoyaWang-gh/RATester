package utils

import (
	"fmt"
	"testing"
)

func TestUint32_2(t *testing.T) {
	tests := []struct {
		name    string
		s       StrTo
		want    uint32
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
			s:       "4294967296",
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

			got, err := tt.s.Uint32()
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint32() got = %v, want %v", got, tt.want)
			}
		})
	}
}
