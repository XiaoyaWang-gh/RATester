package commands

import (
	"fmt"
	"sync"
	"testing"

	"github.com/gohugoio/hugo/common/types"
	"golang.org/x/sync/semaphore"
)

func TestFullRebuild_1(t *testing.T) {
	type fields struct {
		r                  *rootCommand
		confmu             sync.Mutex
		conf               *commonConfig
		s                  *serverCommand
		changeDetector     *fileChangeDetector
		visitedURLs        *types.EvictingStringQueue
		fullRebuildSem     *semaphore.Weighted
		debounce           func(f func())
		onConfigLoaded     func(reloaded bool) error
		fastRenderMode     bool
		showErrorInBrowser bool
		errState           hugoBuilderErrState
	}
	type args struct {
		changeType string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &hugoBuilder{
				r:                  tt.fields.r,
				confmu:             tt.fields.confmu,
				conf:               tt.fields.conf,
				s:                  tt.fields.s,
				changeDetector:     tt.fields.changeDetector,
				visitedURLs:        tt.fields.visitedURLs,
				fullRebuildSem:     tt.fields.fullRebuildSem,
				debounce:           tt.fields.debounce,
				onConfigLoaded:     tt.fields.onConfigLoaded,
				fastRenderMode:     tt.fields.fastRenderMode,
				showErrorInBrowser: tt.fields.showErrorInBrowser,
				errState:           tt.fields.errState,
			}
			c.fullRebuild(tt.args.changeType)
		})
	}
}
