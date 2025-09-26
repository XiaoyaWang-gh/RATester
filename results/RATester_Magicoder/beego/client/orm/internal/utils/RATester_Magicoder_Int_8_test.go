package utils

import (
	"fmt"
	"testing"
)

func TestInt_8(t *testing.T) {
	tests := []struct {
		name    string
		f       StrTo
		want    int
		wantErr bool
	}{
		{
			name:    "Test case 1",
			f:       "123",
			want:    123,
			wantErr: false,
		},
		{
			name:    "Test case 2",
			f:       "abc",
			want:    0,
			wantErr: true,
		},
		{
			name:    "Test case 3",
			f:       "",
			want:    0,
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

			got, err := tt.f.Int()
			if (err != nil) != tt.wantErr {
				t.Errorf("Int() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int() got = %v, want %v", got, tt.want)
			}
		})
	}
}
