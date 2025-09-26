package metrics

import (
	"context"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/types"
)

func TestInitStatsdTicker_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	config := &types.Statsd{}
	config.PushInterval = 10
	config.Address = "localhost:8125"
	config.AddEntryPointsLabels = true
	config.AddRoutersLabels = true
	config.AddServicesLabels = true
	config.Prefix = "traefik"

	report := initStatsdTicker(ctx, config)
	assert.NotNil(t, report)
}
