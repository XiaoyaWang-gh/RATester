package session

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionInit_8(t *testing.T) {
	ctx := context.Background()
	pder := &CookieProvider{}

	type test struct {
		name        string
		maxlifetime int64
		config      string
		wantErr     bool
	}

	tests := []test{
		{
			name:        "valid config",
			maxlifetime: 3600,
			config:      `{"BlockKey": "1234567890123456", "SecurityName": "abcdefghijklmnopqrst"}`,
			wantErr:     false,
		},
		{
			name:        "invalid config",
			maxlifetime: 3600,
			config:      `{"BlockKey": "1234567890123456", "SecurityName": "abcdefghijklmnopqrst", "InvalidField": "invalidValue"}`,
			wantErr:     true,
		},
		{
			name:        "empty config",
			maxlifetime: 3600,
			config:      ``,
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := pder.SessionInit(ctx, tt.maxlifetime, tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("SessionInit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
