package ecs

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/patrickmn/go-cache"
)

func TestLookupTaskDefinitions_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := &awsClient{
		ecs: &ecs.ECS{},
	}

	taskDefArns := map[string]*ecs.Task{
		"arn:aws:ecs:us-west-2:123456789012:task-definition/my-task-def:1": {
			TaskDefinitionArn: aws.String("arn:aws:ecs:us-west-2:123456789012:task-definition/my-task-def:1"),
		},
	}

	existingTaskDefCache = cache.New(5*time.Minute, 10*time.Minute)
	existingTaskDefCache.Set("arn:aws:ecs:us-west-2:123456789012:task-definition/my-task-def:1", &ecs.TaskDefinition{
		TaskDefinitionArn: aws.String("arn:aws:ecs:us-west-2:123456789012:task-definition/my-task-def:1"),
	}, cache.DefaultExpiration)

	taskDef, err := (&Provider{}).lookupTaskDefinitions(ctx, client, taskDefArns)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(taskDef) != 1 {
		t.Errorf("Expected 1 task definition, got %d", len(taskDef))
	}

	_, ok := taskDef["arn:aws:ecs:us-west-2:123456789012:task-definition/my-task-def:1"]
	if !ok {
		t.Errorf("Expected task definition to be in the map")
	}
}
