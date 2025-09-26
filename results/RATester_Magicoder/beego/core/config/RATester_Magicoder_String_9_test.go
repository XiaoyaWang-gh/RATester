package config

import (
	"fmt"
	"testing"
)

func TestString_9(t *testing.T) {
	container := &IniConfigContainer{
		data: map[string]map[string]string{
			"section1": {"key1": "value1", "key2": "value2"},
			"section2": {"key3": "value3", "key4": "value4"},
		},
	}

	tests := []struct {
		name    string
		key     string
		want    string
		wantErr bool
	}{
		{
			name:    "Existing key",
			key:     "section1:key1",
			want:    "value1",
			wantErr: false,
		},
		{
			name:    "Non-existing key",
			key:     "section1:key3",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Invalid key format",
			key:     "section1",
			want:    "",
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

			got, err := container.String(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
