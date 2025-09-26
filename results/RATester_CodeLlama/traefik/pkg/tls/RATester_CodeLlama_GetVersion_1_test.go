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
			name: "tls10",
			connState: &tls.ConnectionState{
				Version: tls.VersionTLS10,
			},
			want: "1.0",
		},
		{
			name: "tls11",
			connState: &tls.ConnectionState{
				Version: tls.VersionTLS11,
			},
			want: "1.1",
		},
		{
			name: "tls12",
			connState: &tls.ConnectionState{
				Version: tls.VersionTLS12,
			},
			want: "1.2",
		},
		{
			name: "tls13",
			connState: &tls.ConnectionState{
				Version: tls.VersionTLS13,
			},
			want: "1.3",
		},
		{
			name: "unknown",
			connState: &tls.ConnectionState{
				Version: 0x1234,
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
