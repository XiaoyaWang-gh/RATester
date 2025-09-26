package server

import (
	"context"
	"fmt"
	"testing"
)

func TestShutdown_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	ctx := context.Background()
	ep := &UDPEntryPoint{}

	// when
	ep.Shutdown(ctx)

	// then
	// TODO
}
