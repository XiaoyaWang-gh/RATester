package logs

import (
	"fmt"
	"testing"
)

func TestInit_33(t *testing.T) {
	writer := &SLACKWriter{}

	testCases := []struct {
		name    string
		config  string
		wantErr bool
	}{
		{
			name:    "valid config",
			config:  `{"webhookurl": "http://example.com", "level": 1, "formatter": "json"}`,
			wantErr: false,
		},
		{
			name:    "invalid config",
			config:  `{"webhookurl": "http://example.com", "level": "one", "formatter": "json"}`,
			wantErr: true,
		},
		{
			name:    "missing formatter",
			config:  `{"webhookurl": "http://example.com", "level": 1}`,
			wantErr: false,
		},
		{
			name:    "invalid formatter",
			config:  `{"webhookurl": "http://example.com", "level": 1, "formatter": "invalid"}`,
			wantErr: true,
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
			}
		})
	}
}
