package docker

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	dockertypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func TestListTasks_1(t *testing.T) {
	type args struct {
		ctx               context.Context
		dockerClient      client.APIClient
		serviceID         string
		serviceDockerData dockerData
		networkMap        map[string]*dockertypes.NetworkResource
		isGlobalSvc       bool
	}
	tests := []struct {
		name    string
		args    args
		want    []dockerData
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

			got, err := listTasks(tt.args.ctx, tt.args.dockerClient, tt.args.serviceID, tt.args.serviceDockerData, tt.args.networkMap, tt.args.isGlobalSvc)
			if (err != nil) != tt.wantErr {
				t.Errorf("listTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}
