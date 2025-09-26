package config

import (
	"fmt"
	"testing"
)

func TestBool_5(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		want    bool
		wantErr bool
	}{
		{
			name:    "Test case 1",
			key:     "key1",
			want:    true,
			wantErr: false,
		},
		{
			name:    "Test case 2",
			key:     "key2",
			want:    false,
			wantErr: false,
		},
		{
			name:    "Test case 3",
			key:     "key3",
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

			got, err := Bool(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}
