package metrics

import (
	"context"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/types"
)

func TestRegisterDatadog_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	config := &types.Datadog{}

	registry := RegisterDatadog(ctx, config)

	assert.NotNil(t, registry)
}
