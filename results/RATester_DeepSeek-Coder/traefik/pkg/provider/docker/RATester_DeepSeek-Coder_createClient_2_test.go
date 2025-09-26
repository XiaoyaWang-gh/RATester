package docker

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/docker/docker/client"
)

func TestCreateClient_2(t *testing.T) {
	type args struct {
		ctx context.Context
		cfg ClientConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *client.Client
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

			got, err := createClient(tt.args.ctx, tt.args.cfg)
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
