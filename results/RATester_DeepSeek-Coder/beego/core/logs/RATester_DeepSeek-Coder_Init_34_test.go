package logs

import (
	"fmt"
	"testing"
)

func TestInit_34(t *testing.T) {
	type testCase struct {
		name    string
		config  string
		wantErr bool
	}

	tests := []testCase{
		{
			name:    "empty config",
			config:  "",
			wantErr: false,
		},
		{
			name:    "valid config",
			config:  `{"formatter": "json", "level": 1, "color": true}`,
			wantErr: false,
		},
		{
			name:    "invalid config",
			config:  `{"formatter": "unknown", "level": 1, "color": true}`,
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

			c := &consoleWriter{}
			err := c.Init(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("consoleWriter.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
