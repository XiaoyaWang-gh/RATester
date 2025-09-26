package health

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestDoTCPCheck_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	monitor := &Monitor{
		addr: "localhost:8080",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := monitor.doTCPCheck(ctx)
	if err != nil {
		t.Errorf("doTCPCheck failed: %v", err)
	}
}
