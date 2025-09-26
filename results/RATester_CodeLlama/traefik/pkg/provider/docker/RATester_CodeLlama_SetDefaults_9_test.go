package docker

import (
	"fmt"
	"testing"
	"time"

	"github.com/alecthomas/assert"
)

func TestSetDefaults_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &SwarmProvider{}
	p.SetDefaults()

	assert.Equal(t, true, p.Watch)
	assert.Equal(t, true, p.ExposedByDefault)
	assert.Equal(t, "unix:///var/run/docker.sock", p.Endpoint)
	assert.Equal(t, 15*time.Second, p.RefreshSeconds)
}
