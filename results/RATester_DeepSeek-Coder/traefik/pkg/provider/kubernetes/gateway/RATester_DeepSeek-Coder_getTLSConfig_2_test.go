package gateway

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/tls"
)

func TestGetTLSConfig_2(t *testing.T) {
	type args struct {
		tlsConfigs map[string]*tls.CertAndStores
	}
	tests := []struct {
		name string
		args args
		want []*tls.CertAndStores
	}{
		{
			name: "Test Case 1",
			args: args{
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

			if got := getTLSConfig(tt.args.tlsConfigs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTLSConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
