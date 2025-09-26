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

	if len(provider.Clusters) != 1 || provider.Clusters[0] != "default" {
		t.Errorf("Expected Clusters to be set to [\"default\"], but got %v", provider.Clusters)
	}

	if provider.AutoDiscoverClusters != false {
		t.Errorf("Expected AutoDiscoverClusters to be set to false, but got %v", provider.AutoDiscoverClusters)
	}

	if provider.HealthyTasksOnly != false {
		t.Errorf("Expected HealthyTasksOnly to be set to false, but got %v", provider.HealthyTasksOnly)
	}

	if provider.ExposedByDefault != true {
		t.Errorf("Expected ExposedByDefault to be set to true, but got %v", provider.ExposedByDefault)
	}

	if provider.RefreshSeconds != 15 {
		t.Errorf("Expected RefreshSeconds to be set to 15, but got %v", provider.RefreshSeconds)
	}

	if provider.DefaultRule != DefaultTemplateRule {
		t.Errorf("Expected DefaultRule to be set to DefaultTemplateRule, but got %v", provider.DefaultRule)
	}
}
