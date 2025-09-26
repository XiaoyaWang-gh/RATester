package metrics

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
)

func TestRegisterPrometheus_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	config := &types.Prometheus{}

	RegisterPrometheus(ctx, config)

	// TODO: test something
}
