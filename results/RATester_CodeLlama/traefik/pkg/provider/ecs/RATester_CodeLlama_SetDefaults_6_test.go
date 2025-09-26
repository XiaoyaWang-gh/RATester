package ecs

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestSetDefaults_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{}
	p.SetDefaults()

	assert.Equal(t, []string{"default"}, p.Clusters)
	assert.Equal(t, false, p.AutoDiscoverClusters)
	assert.Equal(t, false, p.HealthyTasksOnly)
	assert.Equal(t, true, p.ExposedByDefault)
	assert.Equal(t, 15, p.RefreshSeconds)
	assert.Equal(t, DefaultTemplateRule, p.DefaultRule)
}
