package acme

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
)

func TestRegister_1(t *testing.T) {
	type fields struct {
		Configuration *Configuration
		ResolverName  string
		Store         Store
		client        *lego.Client
	}
	type args struct {
		ctx    context.Context
		client *lego.Client
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *registration.Resource
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
				Configuration: tt.fields.Configuration,
				ResolverName:  tt.fields.ResolverName,
				Store:         tt.fields.Store,
				client:        tt.fields.client,
			}
			got, err := p.register(tt.args.ctx, tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Provider.register() = %v, want %v", got, tt.want)
			}
		})
	}
}
