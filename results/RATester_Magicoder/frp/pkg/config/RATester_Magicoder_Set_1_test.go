package config

import (
	"fmt"
	"testing"
)

func TestSet_1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    bool
		wantErr bool
	}{
		{
			name:    "Test case 1",
			input:   "true",
			want:    true,
			wantErr: false,
		},
		{
			name:    "Test case 2",
			input:   "false",
			want:    false,
			wantErr: false,
		},
		{
			name:    "Test case 3",
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

			f := &BoolFuncFlag{}
			err := f.Set(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("BoolFuncFlag.Set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if f.v != tt.want {
				t.Errorf("BoolFuncFlag.Set() = %v, want %v", f.v, tt.want)
			}
		})
	}
}
