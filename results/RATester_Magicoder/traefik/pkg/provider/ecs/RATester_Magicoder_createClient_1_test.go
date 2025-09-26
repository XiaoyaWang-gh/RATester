package ecs

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestCreateClient_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := zerolog.New(os.Stdout)
	provider := &Provider{
		Region: "us-west-2",
	}

	client, err := provider.createClient(logger)
	if err != nil {
		t.Errorf("createClient() error = %v", err)
		return
	}

	if client == nil {
		t.Errorf("createClient() client is nil")
		return
	}

	if client.ecs == nil {
		t.Errorf("createClient() client.ecs is nil")
		return
	}

	if client.ec2 == nil {
		t.Errorf("createClient() client.ec2 is nil")
		return
	}

	if client.ssm == nil {
		t.Errorf("createClient() client.ssm is nil")
		return
	}
}
