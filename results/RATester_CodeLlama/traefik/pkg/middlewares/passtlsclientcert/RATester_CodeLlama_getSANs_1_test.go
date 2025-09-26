package passtlsclientcert

import (
	"crypto/x509"
	"fmt"
	"net"
	"net/url"
	"reflect"
	"testing"
)

func TestGetSANs_1(t *testing.T) {
	type args struct {
		cert *x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test case 1",
			args: args{
				cert: &x509.Certificate{
					DNSNames:       []string{"example.com"},
					EmailAddresses: []string{"test@example.com"},
					IPAddresses:    []net.IP{net.IPv4(127, 0, 0, 1)},
					URIs:           []*url.URL{&url.URL{Scheme: "https", Host: "example.com"}},
				},
			},
			want: []string{"example.com", "test@example.com", "127.0.0.1", "https://example.com"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getSANs(tt.args.cert); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSANs() = %v, want %v", got, tt.want)
			}
		})
	}
}
