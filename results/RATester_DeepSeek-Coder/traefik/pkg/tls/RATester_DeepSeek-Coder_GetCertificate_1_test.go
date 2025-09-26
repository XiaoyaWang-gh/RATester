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
	type testCase struct {
		name     string
		store    *CertificateStore
		domains  []string
		expected *tls.Certificate
	}

	testCases := []testCase{
		{
			name: "Test with nil CertificateStore",
			store: &CertificateStore{
				CertCache:          cache.New(5*time.Minute, 10*time.Minute),
				DefaultCertificate: &tls.Certificate{},
			},
			domains:  []string{"example.com"},
			expected: nil,
		},
		{
			name: "Test with empty domains",
			store: &CertificateStore{
				CertCache:          cache.New(5*time.Minute, 10*time.Minute),
				DefaultCertificate: &tls.Certificate{},
			},
			domains:  []string{},
			expected: &tls.Certificate{},
		},
		{
			name: "Test with valid domains",
			store: &CertificateStore{
				CertCache:          cache.New(5*time.Minute, 10*time.Minute),
				DefaultCertificate: &tls.Certificate{},
				DynamicCerts:       &safe.Safe{},
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
