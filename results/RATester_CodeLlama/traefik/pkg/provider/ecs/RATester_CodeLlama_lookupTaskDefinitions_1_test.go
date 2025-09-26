package ecs

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
)

func TestLookupTaskDefinitions_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	p := &Provider{
		Clusters: []string{"cluster"},
	}
	client := &awsClient{}
	taskDefArns := map[string]*ecs.Task{
		"arn": {
			TaskDefinitionArn: aws.String("arn"),
		},
	}
	taskDef, err := p.lookupTaskDefinitions(ctx, client, taskDefArns)
	if err != nil {
		t.Fatal(err)
	}
	if len(taskDef) != 1 {
		t.Fatalf("expected 1 task definition, got %d", len(taskDef))
	}
	if taskDef["arn"] == nil {
		t.Fatal("expected task definition to be found")
	}
}
