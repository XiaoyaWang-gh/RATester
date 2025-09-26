package config

import (
	"fmt"
	"testing"
)

func TestInt_3(t *testing.T) {
	container := &IniConfigContainer{
		data: map[string]map[string]string{
			"section1": {"key1": "123"},
		},
	}

	tests := []struct {
		name    string
		key     string
		want    int
		wantErr bool
	}{
		{
			name:    "valid int",
			key:     "section1:key1",
			want:    123,
			wantErr: false,
		},
		{
			name:    "invalid int",
			key:     "section1:key1",
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

			got, err := container.Int(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}
