package ingress

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/tls"
)

func TestGetTLSConfig_1(t *testing.T) {
	tests := []struct {
		name       string
		tlsConfigs map[string]*tls.CertAndStores
		want       []*tls.CertAndStores
	}{
		{
			name: "test case 1",
			tlsConfigs: map[string]*tls.CertAndStores{
				"test1": {
					Certificate: tls.Certificate{
						CertFile: "test1.crt",
						KeyFile:  "test1.key",
					},
					Stores: []string{"store1", "store2"},
				},
				"test2": {
					Certificate: tls.Certificate{
						CertFile: "test2.crt",
						KeyFile:  "test2.key",
					},
					Stores: []string{"store3", "store4"},
				},
			},
			want: []*tls.CertAndStores{
				{
					Certificate: tls.Certificate{
						CertFile: "test1.crt",
						KeyFile:  "test1.key",
					},
					Stores: []string{"store1", "store2"},
				},
				{
					Certificate: tls.Certificate{
						CertFile: "test2.crt",
						KeyFile:  "test2.key",
					},
					Stores: []string{"store3", "store4"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getTLSConfig(tt.tlsConfigs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTLSConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
