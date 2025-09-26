package utils

import (
	"fmt"
	"testing"
)

func TestInt8_1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int8
		wantErr bool
	}{
		{
			name:    "valid int8",
			input:   "123",
			want:    123,
			wantErr: false,
		},
		{
			name:    "invalid int8",
			input:   "abc",
			want:    0,
			wantErr: true,
		},
		{
			name:    "out of range int8",
			input:   "300",
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
			got, err := f.Int8()
			if (err != nil) != tt.wantErr {
				t.Errorf("Int8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int8() got = %v, want %v", got, tt.want)
			}
		})
	}
}
