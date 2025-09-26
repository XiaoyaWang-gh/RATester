package ingress

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/safe"
)

func TestThrottleEvents_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	ctx := context.Background()
	throttleDuration := time.Duration(10)
	pool := &safe.Pool{}
	eventsChan := make(chan interface{}, 1)

	// when
	eventsChanBuffered := throttleEvents(ctx, throttleDuration, pool, eventsChan)

	// then
	assert.NotNil(t, eventsChanBuffered)
}
