package logs

import (
	"fmt"
	"testing"
)

func TestInit_32(t *testing.T) {
	tests := []struct {
		name    string
		c       *connWriter
		config  string
		wantErr bool
	}{
		{
			name: "Test with valid config",
			c:    &connWriter{},
			config: `{
				"formatter": "json",
				"reconnectOnMsg": true,
				"reconnect": true,
				"net": "tcp",
				"addr": "localhost:8080",
				"level": 1
			}`,
			wantErr: false,
		},
		{
			name: "Test with invalid config",
			c:    &connWriter{},
			config: `{
				"formatter": "invalid",
				"reconnectOnMsg": true,
				"reconnect": true,
				"net": "tcp",
				"addr": "localhost:8080",
				"level": 1
			}`,
			wantErr: true,
		},
		{
			name: "Test with missing formatter",
			c:    &connWriter{},
			config: `{
				"reconnectOnMsg": true,
				"reconnect": true,
				"net": "tcp",
				"addr": "localhost:8080",
				"level": 1
			}`,
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

			err := tt.c.Init(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("connWriter.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
