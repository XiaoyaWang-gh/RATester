package docker

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/docker/docker/client"
)

func TestGetClientOpts_1(t *testing.T) {
	type args struct {
		ctx context.Context
		cfg ClientConfig
	}
	tests := []struct {
		name    string
		args    args
		want    []client.Opt
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

			got, err := getClientOpts(tt.args.ctx, tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("getClientOpts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getClientOpts() = %v, want %v", got, tt.want)
			}
		})
	}
}
