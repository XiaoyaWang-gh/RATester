package tls

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetCipherSuites_1(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "Test Case 1",
			want: []string{
				"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
				"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
				"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305",
				"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305",
				"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
				"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
				"TLS_RSA_WITH_AES_128_GCM_SHA256",
				"TLS_RSA_WITH_AES_256_GCM_SHA384",
				"TLS_RSA_WITH_AES_128_CBC_SHA",
				"TLS_RSA_WITH_AES_256_CBC_SHA",
				"TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA",
				"TLS_RSA_WITH_3DES_EDE_CBC_SHA",
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getCipherSuites(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCipherSuites() = %v, want %v", got, tt.want)
			}
		})
	}
}
