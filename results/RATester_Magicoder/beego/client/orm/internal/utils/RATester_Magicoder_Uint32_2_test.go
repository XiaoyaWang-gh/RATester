package utils

import (
	"fmt"
	"testing"
)

func TestUint32_2(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    uint32
		wantErr bool
	}{
		{
			name:    "valid uint32",
			input:   "123",
			want:    123,
			wantErr: false,
		},
		{
			name:    "invalid uint32",
			input:   "abc",
			want:    0,
			wantErr: true,
		},
		{
			name:    "out of range uint32",
			input:   "4294967296",
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
			got, err := f.Uint32()
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint32() = %v, want %v", got, tt.want)
			}
		})
	}
}
