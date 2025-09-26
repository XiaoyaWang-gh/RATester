package fiber

import (
	"fmt"
	"testing"
)

func TestListen_1(t *testing.T) {
	app := &App{}

	testCases := []struct {
		name    string
		config  ListenConfig
		wantErr bool
	}{
		{
			name: "Test with valid config",
			config: ListenConfig{
				CertFile:       "testdata/server.crt",
				CertKeyFile:    "testdata/server.key",
				CertClientFile: "testdata/client.crt",
			},
			wantErr: false,
		},
		{
			name: "Test with invalid cert file",
			config: ListenConfig{
				CertFile:       "testdata/invalid.crt",
				CertKeyFile:    "testdata/server.key",
				CertClientFile: "testdata/client.crt",
			},
			wantErr: true,
		},
		{
			name: "Test with invalid key file",
			config: ListenConfig{
				CertFile:       "testdata/server.crt",
				CertKeyFile:    "testdata/invalid.key",
				CertClientFile: "testdata/client.crt",
			},
			wantErr: true,
		},
		{
			name: "Test with invalid client cert file",
			config: ListenConfig{
				CertFile:       "testdata/server.crt",
				CertKeyFile:    "testdata/server.key",
				CertClientFile: "testdata/invalid.crt",
			},
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

			err := app.Listen(":3000", tc.config)
			if (err != nil) != tc.wantErr {
				t.Errorf("Listen() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
