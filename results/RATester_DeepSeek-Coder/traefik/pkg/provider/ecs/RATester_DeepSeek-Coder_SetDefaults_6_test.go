package ecs

import (
	"fmt"
	"testing"
)

func TestSetDefaults_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := &Provider{}
	provider.SetDefaults()

	if provider.Clusters[0] != "default" {
		t.Errorf("Expected Clusters to be 'default', got %s", provider.Clusters[0])
	}
	if provider.AutoDiscoverClusters != false {
		t.Errorf("Expected AutoDiscoverClusters to be false, got %v", provider.AutoDiscoverClusters)
	}
	if provider.HealthyTasksOnly != false {
		t.Errorf("Expected HealthyTasksOnly to be false, got %v", provider.HealthyTasksOnly)
	}
	if provider.ExposedByDefault != true {
		t.Errorf("Expected ExposedByDefault to be true, got %v", provider.ExposedByDefault)
	}
	if provider.RefreshSeconds != 15 {
		t.Errorf("Expected RefreshSeconds to be 15, got %v", provider.RefreshSeconds)
	}
}
