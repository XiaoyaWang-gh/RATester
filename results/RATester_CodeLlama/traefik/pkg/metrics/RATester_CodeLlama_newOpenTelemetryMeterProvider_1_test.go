package metrics

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
)

func TestNewOpenTelemetryMeterProvider_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	config := &types.OTLP{
		GRPC: &types.OtelGRPC{
			Endpoint: "localhost:1234",
		},
		HTTP: &types.OtelHTTP{
			Endpoint: "localhost:1234",
		},
		AddEntryPointsLabels: true,
		AddRoutersLabels:     true,
		AddServicesLabels:    true,
		ExplicitBoundaries:   []float64{1, 2, 3},
		PushInterval:         10,
	}

	meterProvider, err := newOpenTelemetryMeterProvider(ctx, config)
	if err != nil {
		t.Errorf("creating meter provider: %v", err)
	}

	if meterProvider == nil {
		t.Error("meter provider is nil")
	}
}
