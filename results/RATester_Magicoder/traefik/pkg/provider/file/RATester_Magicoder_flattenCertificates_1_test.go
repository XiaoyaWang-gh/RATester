package file

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/tls"
)

func TestflattenCertificates_1(t *testing.T) {
	type args struct {
		ctx       context.Context
		tlsConfig *dynamic.TLSConfiguration
	}
	tests := []struct {
		name string
		args args
		want []*tls.CertAndStores
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

			if got := flattenCertificates(tt.args.ctx, tt.args.tlsConfig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flattenCertificates() = %v, want %v", got, tt.want)
			}
		})
	}
}
