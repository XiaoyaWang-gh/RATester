package visitor

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestUpdateAll_1(t *testing.T) {
	type fields struct {
		clientCfg               *v1.ClientCommonConfig
		cfgs                    map[string]v1.VisitorConfigurer
		visitors                map[string]Visitor
		helper                  Helper
		checkInterval           time.Duration
		keepVisitorsRunningOnce sync.Once
		mu                      sync.RWMutex
		ctx                     context.Context
		stopCh                  chan struct{}
	}
	type args struct {
		cfgs []v1.VisitorConfigurer
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

			m := &Manager{
				clientCfg:               tt.fields.clientCfg,
				cfgs:                    tt.fields.cfgs,
				visitors:                tt.fields.visitors,
				helper:                  tt.fields.helper,
				checkInterval:           tt.fields.checkInterval,
				keepVisitorsRunningOnce: tt.fields.keepVisitorsRunningOnce,
				mu:                      tt.fields.mu,
				ctx:                     tt.fields.ctx,
				stopCh:                  tt.fields.stopCh,
			}
			m.UpdateAll(tt.args.cfgs)
		})
	}
}
