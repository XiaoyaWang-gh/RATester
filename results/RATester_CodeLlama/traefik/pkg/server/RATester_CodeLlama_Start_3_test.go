package server

import (
	"context"
	"fmt"
	"testing"
)

func TestStart_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	ep := &UDPEntryPoint{}
	ep.Start(ctx)
}
