package tls

import (
	"crypto/tls"
	"fmt"
	"testing"
)

func TestGetCipherName_1(t *testing.T) {
	tests := []struct {
		name string
		conn *tls.ConnectionState
		want string
	}{
		{
			name: "Test with known cipher suite",
			conn: &tls.ConnectionState{
				CipherSuite: tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			},
			want: "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
		},
		{
			name: "Test with unknown cipher suite",
			conn: &tls.ConnectionState{
				CipherSuite: 0xFFFF,
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

			if got := GetCipherName(tt.conn); got != tt.want {
				t.Errorf("GetCipherName() = %v, want %v", got, tt.want)
			}
		})
	}
}
