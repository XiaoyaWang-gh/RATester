package consulcatalog

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	traefiktls "github.com/traefik/traefik/v3/pkg/tls"
	"github.com/traefik/traefik/v3/pkg/types"
)

func TestTcpServersTransport_1(t *testing.T) {
	type testCase struct {
		name     string
		input    itemData
		expected *dynamic.TCPServersTransport
	}

	testCases := []testCase{
		{
			name: "test case 1",
			input: itemData{
				Namespace:  "test",
				Datacenter: "dc1",
				Name:       "service1",
			},
			expected: &dynamic.TCPServersTransport{
				TLS: &dynamic.TLSClientConfig{
					ServerName:         "test-dc1-service1",
					InsecureSkipVerify: true,
					RootCAs:            []types.FileOrContent{},
					Certificates:       traefiktls.Certificates{},
					PeerCertURI:        "spiffe:///ns/test/dc/dc1/svc/service1",
				},
			},
		},
		// add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &connectCert{}
			result := c.tcpServersTransport(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
