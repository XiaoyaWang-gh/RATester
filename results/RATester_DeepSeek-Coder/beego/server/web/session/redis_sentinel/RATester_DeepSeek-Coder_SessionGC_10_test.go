package redis_sentinel

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionGC_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	rp := &Provider{
		maxlifetime: 10,
		SavePath:    "/tmp",
		Poolsize:    10,
		Password:    "password",
	}

	rp.SessionGC(ctx)

	// Add assertions here to verify the behavior of SessionGC
}
