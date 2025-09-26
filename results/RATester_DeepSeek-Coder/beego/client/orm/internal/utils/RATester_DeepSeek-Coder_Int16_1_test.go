package utils

import (
	"fmt"
	"testing"
)

func TestInt16_1(t *testing.T) {
	tests := []struct {
		name    string
		f       StrTo
		want    int16
		wantErr bool
	}{
		{
			name:    "valid",
			f:       StrTo("32767"),
			want:    32767,
			wantErr: false,
		},
		{
			name:    "invalid",
			f:       StrTo("32768"),
			want:    0,
			wantErr: true,
		},
		{
			name:    "empty",
			f:       StrTo(""),
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

			got, err := tt.f.Int16()
			if (err != nil) != tt.wantErr {
				t.Errorf("Int16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int16() got = %v, want %v", got, tt.want)
			}
		})
	}
}
