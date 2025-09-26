package visitor

import (
	"context"
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"golang.org/x/time/rate"
)

func TestProcessTunnelStartEvents_1(t *testing.T) {
	type fields struct {
		BaseVisitor   *BaseVisitor
		session       TunnelSession
		startTunnelCh chan struct{}
		retryLimiter  *rate.Limiter
		cancel        context.CancelFunc
		cfg           *v1.XTCPVisitorConfig
	}
	tests := []struct {
		name   string
		fields fields
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

			sv := &XTCPVisitor{
				BaseVisitor:   tt.fields.BaseVisitor,
				session:       tt.fields.session,
				startTunnelCh: tt.fields.startTunnelCh,
				retryLimiter:  tt.fields.retryLimiter,
				cancel:        tt.fields.cancel,
				cfg:           tt.fields.cfg,
			}
			sv.processTunnelStartEvents()
		})
	}
}
