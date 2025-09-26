package utils

import (
	"fmt"
	"testing"
)

func TestInt64_8(t *testing.T) {
	tests := []struct {
		name    string
		f       StrTo
		want    int64
		wantErr bool
	}{
		{
			name:    "Test with valid string",
			f:       StrTo("123"),
			want:    123,
			wantErr: false,
		},
		{
			name:    "Test with invalid string",
			f:       StrTo("abc"),
			want:    0,
			wantErr: true,
		},
		{
			name:    "Test with large number",
			f:       StrTo("9223372036854775807"),
			want:    9223372036854775807,
			wantErr: false,
		},
		{
			name:    "Test with negative number",
			f:       StrTo("-123"),
			want:    -123,
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

			got, err := tt.f.Int64()
			if (err != nil) != tt.wantErr {
				t.Errorf("Int64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int64() got = %v, want %v", got, tt.want)
			}
		})
	}
}
