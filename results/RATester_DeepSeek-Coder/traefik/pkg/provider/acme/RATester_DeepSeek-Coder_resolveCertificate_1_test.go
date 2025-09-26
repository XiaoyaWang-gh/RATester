package acme

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/go-acme/lego/v4/certificate"
	"github.com/traefik/traefik/v3/pkg/types"
)

func TestResolveCertificate_1(t *testing.T) {
	type args struct {
		ctx      context.Context
		domain   types.Domain
		tlsStore string
	}
	tests := []struct {
		name    string
		args    args
		want    types.Domain
		want1   *certificate.Resource
		wantErr bool
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

			p := &Provider{}
			got, got1, err := p.resolveCertificate(tt.args.ctx, tt.args.domain, tt.args.tlsStore)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.resolveCertificate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Provider.resolveCertificate() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Provider.resolveCertificate() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
