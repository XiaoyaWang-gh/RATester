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
		// TODO: Add test cases.
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
