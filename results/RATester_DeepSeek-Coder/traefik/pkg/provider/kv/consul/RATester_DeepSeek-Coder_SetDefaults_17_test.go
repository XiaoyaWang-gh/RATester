package consul

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/provider/kv"
	"github.com/traefik/traefik/v3/pkg/types"
	"gotest.tools/assert"
)

func TestSetDefaults_17(t *testing.T) {
	testCases := []struct {
		desc     string
		builder  *ProviderBuilder
		expected *ProviderBuilder
	}{
		{
			desc: "should set defaults correctly",
			builder: &ProviderBuilder{
				Provider: kv.Provider{
					RootKey:   "test",
					Endpoints: []string{"localhost:8500"},
				},
				Token: "test-token",
				TLS: &types.ClientTLS{
					InsecureSkipVerify: true,
				},
				Namespaces: []string{"test-namespace"},
			},
			expected: &ProviderBuilder{
				Provider: kv.Provider{
					RootKey:   "test",
					Endpoints: []string{"127.0.0.1:8500"},
				},
				Token: "test-token",
				TLS: &types.ClientTLS{
					InsecureSkipVerify: true,
				},
				Namespaces: []string{"test-namespace"},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			test.builder.SetDefaults()

			assert.Equal(t, test.expected, test.builder)
		})
	}
}
