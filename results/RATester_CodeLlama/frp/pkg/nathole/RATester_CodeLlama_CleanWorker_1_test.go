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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c := &Controller{
		clientCfgs: make(map[string]*ClientCfg),
		sessions:   make(map[string]*Session),
		analyzer:   &Analyzer{records: make(map[string]*MakeHoleRecords)},
	}
	go c.CleanWorker(ctx)
	time.Sleep(time.Second)
	cancel()
}
