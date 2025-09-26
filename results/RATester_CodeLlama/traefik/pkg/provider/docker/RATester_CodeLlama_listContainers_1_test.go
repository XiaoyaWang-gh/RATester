package docker

import (
	"context"
	"fmt"
	"testing"

	"github.com/docker/docker/client"
)

func TestListContainers_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	dockerClient := &client.Client{}
	p := &Provider{}
	p.listContainers(ctx, dockerClient)
}
