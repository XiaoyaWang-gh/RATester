package consulcatalog

import (
	"fmt"
	"testing"
	"time"

	ptypes "github.com/traefik/paerser/types"
	"github.com/zeebo/assert"
)

func TestSetDefaults_18(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := &Configuration{}
	config.SetDefaults()

	assert.Equal(t, &EndpointConfig{}, config.Endpoint)
	assert.Equal(t, ptypes.Duration(15*time.Second), config.RefreshInterval)
	assert.Equal(t, "traefik", config.Prefix)
	assert.True(t, config.ExposedByDefault)
	assert.Equal(t, defaultTemplateRule, config.DefaultRule)
	assert.Equal(t, "traefik", config.ServiceName)
	assert.Equal(t, defaultStrictChecks(), config.StrictChecks)
}
