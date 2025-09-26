package consulcatalog

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hashicorp/consul/api"
)

func TestCreateClient_4(t *testing.T) {
	type args struct {
		namespace string
		endpoint  *EndpointConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *api.Client
		wantErr bool
	}{
		{
			name: "Test with valid endpoint",
			args: args{
				namespace: "default",
				endpoint: &EndpointConfig{
					Address: "localhost:8500",
					Scheme:  "http",
				},
			},
			want:    &api.Client{},
			wantErr: false,
		},
		{
			name: "Test with invalid endpoint",
			args: args{
				namespace: "default",
				endpoint: &EndpointConfig{
					Address: "invalid:8500",
					Scheme:  "http",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := createClient(tt.args.namespace, tt.args.endpoint)
			if (err != nil) != tt.wantErr {
				t.Errorf("createClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
