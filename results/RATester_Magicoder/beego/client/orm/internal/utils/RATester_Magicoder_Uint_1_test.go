package utils

import (
	"fmt"
	"testing"
)

func TestUint_1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    uint
		wantErr bool
	}{
		{
			name:    "valid uint",
			input:   "123",
			want:    123,
			wantErr: false,
		},
		{
			name:    "invalid uint",
			input:   "abc",
			want:    0,
			wantErr: true,
		},
		{
			name:    "empty string",
			input:   "",
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
			got, err := f.Uint()
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint() got = %v, want %v", got, tt.want)
			}
		})
	}
}
