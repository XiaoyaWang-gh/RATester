package utils

import (
	"fmt"
	"testing"
)

func TestInt8_1(t *testing.T) {
	tests := []struct {
		name    string
		s       StrTo
		want    int8
		wantErr bool
	}{
		{
			name:    "TestInt8_1",
			s:       "123",
			want:    123,
			wantErr: false,
		},
		{
			name:    "TestInt8_2",
			s:       "-123",
			want:    -123,
			wantErr: false,
		},
		{
			name:    "TestInt8_3",
			s:       "256",
			want:    0,
			wantErr: true,
		},
		{
			name:    "TestInt8_4",
			s:       "abc",
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

			got, err := tt.s.Int8()
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
