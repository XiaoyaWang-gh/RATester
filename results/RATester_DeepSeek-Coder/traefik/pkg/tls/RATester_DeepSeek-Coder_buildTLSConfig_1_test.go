package tls

import (
	"fmt"
	"testing"
)

func TestBuildTLSConfig_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		tlsOpt  Options
		wantErr bool
	}{
		{
			name: "valid options",
			tlsOpt: Options{
				MinVersion: "VersionTLS12",
				CipherSuites: []string{
					"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
				},
			},
			wantErr: false,
		},
		{
			name: "invalid minVersion",
			tlsOpt: Options{
				MinVersion: "InvalidVersion",
			},
			wantErr: true,
		},
		{
			name: "invalid cipherSuite",
			tlsOpt: Options{
				CipherSuites: []string{
					"InvalidCipherSuite",
				},
			},
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

			_, err := buildTLSConfig(tt.tlsOpt)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildTLSConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
