package nathole

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCleanWorker_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		clientCfgs: make(map[string]*ClientCfg),
		sessions:   make(map[string]*Session),
		analyzer:   &Analyzer{records: make(map[string]*MakeHoleRecords)},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go ctrl.CleanWorker(ctx)

	time.Sleep(time.Second)
	cancel()

	// check if CleanWorker has exited
	select {
	case <-ctx.Done():
		return
	default:
		t.Errorf("CleanWorker did not exit")
	}
}
