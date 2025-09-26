package consulcatalog

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/hashicorp/consul/api"
)

func TestFetchService_2(t *testing.T) {
	type args struct {
		ctx            context.Context
		name           string
		connectEnabled bool
	}
	tests := []struct {
		name    string
		p       *Provider
		args    args
		want    []*api.ServiceEntry
		want1   map[string]string
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

			got, got1, err := tt.p.fetchService(tt.args.ctx, tt.args.name, tt.args.connectEnabled)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.fetchService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Provider.fetchService() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Provider.fetchService() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
