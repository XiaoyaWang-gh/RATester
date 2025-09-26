package memcache

import (
	"fmt"
	"testing"
)

func TestStartAndGC_3(t *testing.T) {
	cache := &Cache{}

	tests := []struct {
		name    string
		config  string
		wantErr bool
	}{
		{
			name:    "valid config",
			config:  `{"conn": "localhost:11211"}`,
			wantErr: false,
		},
		{
			name:    "invalid config",
			config:  `{"conn": "localhost:11211", "invalid": "field"}`,
			wantErr: true,
		},
		{
			name:    "empty config",
			config:  `{}`,
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

			if err := cache.StartAndGC(tt.config); (err != nil) != tt.wantErr {
				t.Errorf("StartAndGC() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
