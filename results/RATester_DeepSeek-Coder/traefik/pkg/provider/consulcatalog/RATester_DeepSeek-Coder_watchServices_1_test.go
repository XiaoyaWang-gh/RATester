package consulcatalog

import (
	"context"
	"fmt"
	"testing"
	"text/template"

	"github.com/hashicorp/consul/api"
)

func TestWatchServices_1(t *testing.T) {
	type fields struct {
		Configuration     Configuration
		name              string
		namespace         string
		client            *api.Client
		defaultRuleTpl    *template.Template
		certChan          chan *connectCert
		watchServicesChan chan struct{}
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
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

			p := &Provider{
				Configuration:     tt.fields.Configuration,
				name:              tt.fields.name,
				namespace:         tt.fields.namespace,
				client:            tt.fields.client,
				defaultRuleTpl:    tt.fields.defaultRuleTpl,
				certChan:          tt.fields.certChan,
				watchServicesChan: tt.fields.watchServicesChan,
			}
			if err := p.watchServices(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Provider.watchServices() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
