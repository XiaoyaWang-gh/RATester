package acme

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/go-acme/lego/v4/certificate"
)

func TestResolveDefaultCertificate_1(t *testing.T) {
	type args struct {
		ctx     context.Context
		domains []string
	}
	tests := []struct {
		name    string
		args    args
		want    *certificate.Resource
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
			got, err := p.resolveDefaultCertificate(tt.args.ctx, tt.args.domains)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.resolveDefaultCertificate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Provider.resolveDefaultCertificate() = %v, want %v", got, tt.want)
			}
		})
	}
}
