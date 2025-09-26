package nomad

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hashicorp/nomad/api"
)

func TestCreateClient_3(t *testing.T) {
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
		// TODO: Add test cases.
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
