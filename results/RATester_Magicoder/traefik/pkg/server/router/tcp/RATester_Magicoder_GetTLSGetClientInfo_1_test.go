package tcp

import (
	"crypto/tls"
	"fmt"
	"testing"
)

func TestGetTLSGetClientInfo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	router := &Router{
		hostHTTPTLSConfig: map[string]*tls.Config{
			"example.com": {
				MinVersion: tls.VersionTLS12,
			},
		},
		httpsTLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS13,
		},
	}

	getClientInfo := router.GetTLSGetClientInfo()

	tests := []struct {
		serverName string
		minVersion uint16
	}{
		{"example.com", tls.VersionTLS12},
		{"unknown.com", tls.VersionTLS13},
	}

	for _, test := range tests {
		info := &tls.ClientHelloInfo{
			ServerName: test.serverName,
		}

		config, err := getClientInfo(info)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if config.MinVersion != test.minVersion {
			t.Errorf("Expected MinVersion %d, got %d", test.minVersion, config.MinVersion)
		}
	}
}
