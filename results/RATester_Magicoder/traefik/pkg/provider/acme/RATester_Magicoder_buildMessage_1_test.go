package acme

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestBuildMessage_1(t *testing.T) {
	type fields struct {
		Configuration *Configuration
		ResolverName  string
		Store         Store
		certificates  []*CertAndStore
	}
	tests := []struct {
		name   string
		fields fields
		want   dynamic.Message
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
				Configuration: tt.fields.Configuration,
				ResolverName:  tt.fields.ResolverName,
				Store:         tt.fields.Store,
				certificates:  tt.fields.certificates,
			}
			if got := p.buildMessage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Provider.buildMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
