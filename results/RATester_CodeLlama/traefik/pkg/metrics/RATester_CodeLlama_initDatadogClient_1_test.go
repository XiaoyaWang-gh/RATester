package metrics

import (
	"context"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/types"
)

func TestInitDatadogClient_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	config := &types.Datadog{}

	initDatadogClient(ctx, config)

	assert.NotNil(t, ctx)
	assert.NotNil(t, datadogClient)
}
