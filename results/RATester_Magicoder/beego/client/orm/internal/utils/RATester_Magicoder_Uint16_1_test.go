package utils

import (
	"fmt"
	"testing"
)

func TestUint16_1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    uint16
		wantErr bool
	}{
		{
			name:    "valid uint16",
			input:   "65535",
			want:    65535,
			wantErr: false,
		},
		{
			name:    "invalid uint16",
			input:   "65536",
			want:    0,
			wantErr: true,
		},
		{
			name:    "empty string",
			input:   "",
			want:    0,
			wantErr: true,
		},
		{
			name:    "non-numeric string",
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
			got, err := f.Uint16()
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
