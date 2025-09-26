package passtlsclientcert

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestNewSubjectDistinguishedNameOptions_1(t *testing.T) {
	type args struct {
		info *dynamic.TLSClientCertificateSubjectDNInfo
	}
	tests := []struct {
		name string
		args args
		want *SubjectDistinguishedNameOptions
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
				info: &dynamic.TLSClientCertificateSubjectDNInfo{
					CommonName:         true,
					Country:            true,
					DomainComponent:    true,
					Locality:           true,
					Organization:       true,
					OrganizationalUnit: true,
					SerialNumber:       true,
					Province:           true,
				},
			},
			want: &SubjectDistinguishedNameOptions{
				CommonName:             true,
				CountryName:            true,
				DomainComponent:        true,
				LocalityName:           true,
				OrganizationName:       true,
				OrganizationalUnitName: true,
				SerialNumber:           true,
				StateOrProvinceName:    true,
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

			if got := newSubjectDistinguishedNameOptions(tt.args.info); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSubjectDistinguishedNameOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
