package acme

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	"github.com/go-acme/lego/v4/lego"
)

func TestGetClient_1(t *testing.T) {
	type fields struct {
		Configuration *Configuration
		ResolverName  string
		Store         Store
		client        *lego.Client
		clientMutex   sync.Mutex
	}
	tests := []struct {
		name    string
		fields  fields
		want    *lego.Client
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
				clientMutex:   tt.fields.clientMutex,
			}
			got, err := p.getClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("getClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}
