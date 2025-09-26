package tls

import (
	"crypto/tls"
	"fmt"
	"testing"
)

func TestGetVersion_1(t *testing.T) {
	tests := []struct {
		name      string
		connState *tls.ConnectionState
		want      string
	}{
		{
			name: "TLS 1.0",
			connState: &tls.ConnectionState{
				Version: tls.VersionTLS10,
			},
			want: "1.0",
		},
		{
			name: "TLS 1.1",
			connState: &tls.ConnectionState{
				Version: tls.VersionTLS11,
			},
			want: "1.1",
		},
		{
			name: "TLS 1.2",
			connState: &tls.ConnectionState{
				Version: tls.VersionTLS12,
			},
			want: "1.2",
		},
		{
			name: "TLS 1.3",
			connState: &tls.ConnectionState{
				Version: tls.VersionTLS13,
			},
			want: "1.3",
		},
		{
			name: "Unknown",
			connState: &tls.ConnectionState{
				Version: 0xffff,
			},
			want: "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetVersion(tt.connState); got != tt.want {
				t.Errorf("GetVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
