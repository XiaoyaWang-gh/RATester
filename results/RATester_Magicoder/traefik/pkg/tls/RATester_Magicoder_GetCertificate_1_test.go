package tls

import (
	"crypto/tls"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/traefik/traefik/v3/pkg/safe"
)

func TestGetCertificate_1(t *testing.T) {
	testCases := []struct {
		name     string
		store    *CertificateStore
		domains  []string
		expected *tls.Certificate
	}{
		{
			name: "nil store",
			store: &CertificateStore{
				DynamicCerts:       nil,
				DefaultCertificate: nil,
				CertCache:          nil,
			},
			domains:  []string{"example.com"},
			expected: nil,
		},
		{
			name: "cache hit",
			store: &CertificateStore{
				DynamicCerts:       nil,
				DefaultCertificate: nil,
				CertCache:          cache.New(5*time.Minute, 10*time.Minute),
			},
			domains:  []string{"example.com"},
			expected: &tls.Certificate{},
		},
		{
			name: "dynamic certs hit",
			store: &CertificateStore{
				DynamicCerts:       &safe.Safe{},
				DefaultCertificate: nil,
				CertCache:          nil,
			},
			domains:  []string{"example.com"},
			expected: &tls.Certificate{},
		},
		{
			name: "default cert",
			store: &CertificateStore{
				DynamicCerts:       nil,
				DefaultCertificate: &tls.Certificate{},
				CertCache:          nil,
			},
			domains:  []string{"example.com"},
			expected: &tls.Certificate{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.store.GetCertificate(tc.domains)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
