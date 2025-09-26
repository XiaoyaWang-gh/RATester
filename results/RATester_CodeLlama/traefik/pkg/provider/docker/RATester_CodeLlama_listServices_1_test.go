package docker

import (
	"context"
	"fmt"
	"testing"

	"github.com/docker/docker/client"
)

func TestListServices_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	dockerClient := &client.Client{}
	p := &SwarmProvider{}

	dockerDataList, err := p.listServices(ctx, dockerClient)
	if err != nil {
		t.Errorf("listServices returned an error: %v", err)
	}

	if len(dockerDataList) != 0 {
		t.Errorf("listServices returned an unexpected number of services: %v", len(dockerDataList))
	}
}
