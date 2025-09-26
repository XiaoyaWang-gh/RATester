package tcp

import (
	"crypto/tls"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/tcp"
)

func TestAcmeTLSALPNHandler_1(t *testing.T) {
	tests := []struct {
		name           string
		httpsTLSConfig *tls.Config
		expected       tcp.Handler
	}{
		{
			name:           "case 1: httpsTLSConfig is nil",
			httpsTLSConfig: nil,
			expected:       &brokenTLSRouter{},
		},
		{
			name: "case 2: httpsTLSConfig is not nil",
			httpsTLSConfig: &tls.Config{
				NextProtos: []string{"h2", "http/1.1"},
			},
			expected: tcp.HandlerFunc(func(conn tcp.WriteCloser) {
				_ = tls.Server(conn, &tls.Config{
					NextProtos: []string{"h2", "http/1.1"},
				}).Handshake()
			}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			r := &Router{
				httpsTLSConfig: test.httpsTLSConfig,
			}

			result := r.acmeTLSALPNHandler()

			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
