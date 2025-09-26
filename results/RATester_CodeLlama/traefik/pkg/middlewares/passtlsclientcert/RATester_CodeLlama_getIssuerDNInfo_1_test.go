package passtlsclientcert

import (
	"context"
	"crypto/x509/pkix"
	"fmt"
	"testing"
)

func TestGetIssuerDNInfo_1(t *testing.T) {
	type args struct {
		ctx     context.Context
		options *IssuerDistinguishedNameOptions
		cs      *pkix.Name
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				ctx:     context.Background(),
				options: &IssuerDistinguishedNameOptions{},
				cs: &pkix.Name{
					Country:            []string{"US"},
					Organization:       []string{"Acme Co"},
					OrganizationalUnit: []string{"OU"},
					Locality:           []string{"Mountain View"},
					Province:           []string{"CA"},
					StreetAddress:      []string{"1600 Amphitheatre Pkwy"},
					PostalCode:         []string{"94043"},
					SerialNumber:       "1234",
					CommonName:         "www.example.com",
				},
			},
			want: "C=US,O=Acme Co,OU=OU,L=Mountain View,ST=CA,SN=1234,CN=www.example.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getIssuerDNInfo(tt.args.ctx, tt.args.options, tt.args.cs); got != tt.want {
				t.Errorf("getIssuerDNInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
