package utils

import (
	"fmt"
	"testing"
)

func TestUint64_2(t *testing.T) {
	tests := []struct {
		name    string
		s       StrTo
		want    uint64
		wantErr bool
	}{
		{
			name:    "valid uint64",
			s:       StrTo("18446744073709551615"),
			want:    18446744073709551615,
			wantErr: false,
		},
		{
			name:    "invalid uint64",
			s:       StrTo("18446744073709551616"),
			want:    0,
			wantErr: true,
		},
		{
			name:    "negative uint64",
			s:       StrTo("-1"),
			want:    0,
			wantErr: true,
		},
		{
			name:    "float uint64",
			s:       StrTo("18446744073709551615.5"),
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

			got, err := tt.s.Uint64()
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint64() got = %v, want %v", got, tt.want)
			}
		})
	}
}
