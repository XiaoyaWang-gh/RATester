package hugolib

import (
	"context"
	"fmt"
	"testing"

	"github.com/bep/logg"
	"github.com/fsnotify/fsnotify"
)

func TestProcessPartialFileEvents_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var h HugoSites
	var ctx context.Context
	var l logg.LevelLogger
	var config *BuildCfg
	var init func(config *BuildCfg) error
	var events []fsnotify.Event
	h.processPartialFileEvents(ctx, l, config, init, events)
}
