package logs

import (
	"fmt"
	"testing"
)

func TestInit_38(t *testing.T) {
	writer := &SMTPWriter{}

	testCases := []struct {
		name    string
		config  string
		wantErr bool
	}{
		{
			name:    "valid config",
			config:  `{"username": "user", "password": "pass", "host": "localhost", "subject": "subject", "fromAddress": "from", "sendTos": ["to1", "to2"], "level": 1, "insecureSkipVerify": true, "formatter": "json"}`,
			wantErr: false,
		},
		{
			name:    "invalid config",
			config:  `{"username": "user", "password": "pass", "host": "localhost", "subject": "subject", "fromAddress": "from", "sendTos": ["to1", "to2"], "level": 1, "insecureSkipVerify": true, "formatter": "invalid"}`,
			wantErr: true,
		},
		{
			name:    "empty config",
			config:  `{}`,
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := writer.Init(tc.config)
			if (err != nil) != tc.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
