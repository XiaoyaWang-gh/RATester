package cache

import (
	"fmt"
	"testing"
)

func TestStartAndGC_1(t *testing.T) {
	bc := &MemoryCache{
		items: make(map[string]*MemoryItem),
	}

	tests := []struct {
		name    string
		config  string
		wantErr bool
	}{
		{
			name:    "valid config",
			config:  `{"interval": 10}`,
			wantErr: false,
		},
		{
			name:    "invalid config",
			config:  `{"interval": "ten"}`,
			wantErr: true,
		},
		{
			name:    "empty config",
			config:  `{}`,
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

			if err := bc.StartAndGC(tt.config); (err != nil) != tt.wantErr {
				t.Errorf("StartAndGC() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
