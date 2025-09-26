package utils

import (
	"fmt"
	"testing"
)

func TestBool_8(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    bool
		wantErr bool
	}{
		{
			name:    "TestBool_True",
			input:   "true",
			want:    true,
			wantErr: false,
		},
		{
			name:    "TestBool_False",
			input:   "false",
			want:    false,
			wantErr: false,
		},
		{
			name:    "TestBool_Invalid",
			input:   "invalid",
			want:    false,
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
			got, err := f.Bool()
			if (err != nil) != tt.wantErr {
				t.Errorf("StrTo.Bool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StrTo.Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}
