package utils

import (
	"fmt"
	"testing"
)

func TestInt64_8(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int64
		wantErr bool
	}{
		{
			name:    "valid int64",
			input:   "1234567890",
			want:    1234567890,
			wantErr: false,
		},
		{
			name:    "invalid int64",
			input:   "abc",
			want:    0,
			wantErr: true,
		},
		{
			name:    "large int64",
			input:   "9223372036854775807",
			want:    9223372036854775807,
			wantErr: false,
		},
		{
			name:    "negative int64",
			input:   "-1234567890",
			want:    -1234567890,
			wantErr: false,
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
			got, err := f.Int64()
			if (err != nil) != tt.wantErr {
				t.Errorf("Int64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}
