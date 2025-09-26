package ecs

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func TestListInstances_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	client := &awsClient{
		ecs: &ecs.ECS{},
		ec2: &ec2.EC2{},
		ssm: &ssm.SSM{},
	}
	p := &Provider{
		AutoDiscoverClusters: true,
	}

	instances, err := p.listInstances(ctx, client)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(instances) == 0 {
		t.Errorf("Expected at least one instance, got %v", len(instances))
	}
}
