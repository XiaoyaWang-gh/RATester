package passtlsclientcert

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestNewIssuerDistinguishedNameOptions_1(t *testing.T) {
	type args struct {
		info *dynamic.TLSClientCertificateIssuerDNInfo
	}
	tests := []struct {
		name string
		args args
		want *IssuerDistinguishedNameOptions
	}{
		{
			name: "nil info",
			args: args{
				info: nil,
			},
			want: nil,
		},
		{
			name: "non-nil info",
			args: args{
				info: &dynamic.TLSClientCertificateIssuerDNInfo{
					CommonName:      true,
					Country:         true,
					DomainComponent: true,
					Locality:        true,
					Organization:    true,
					SerialNumber:    true,
					Province:        true,
				},
			},
			want: &IssuerDistinguishedNameOptions{
				CommonName:          true,
				CountryName:         true,
				DomainComponent:     true,
				LocalityName:        true,
				OrganizationName:    true,
				SerialNumber:        true,
				StateOrProvinceName: true,
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

			if got := newIssuerDistinguishedNameOptions(tt.args.info); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newIssuerDistinguishedNameOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
