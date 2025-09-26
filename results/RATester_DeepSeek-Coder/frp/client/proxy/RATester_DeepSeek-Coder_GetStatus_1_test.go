package proxy

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/fatedier/frp/client/event"
	"github.com/fatedier/frp/client/health"
	"github.com/fatedier/frp/pkg/transport"
	"github.com/fatedier/frp/pkg/util/xlog"
)

func TestGetStatus_1(t *testing.T) {
	type fields struct {
		WorkingStatus    WorkingStatus
		pxy              Proxy
		monitor          *health.Monitor
		handler          event.Handler
		msgTransporter   transport.MessageTransporter
		health           uint32
		lastSendStartMsg time.Time
		lastStartErr     time.Time
		closeCh          chan struct{}
		healthNotifyCh   chan struct{}
		mu               sync.RWMutex
		xl               *xlog.Logger
		ctx              context.Context
	}
	tests := []struct {
		name   string
		fields fields
		want   *WorkingStatus
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

			pw := &Wrapper{
				WorkingStatus:    tt.fields.WorkingStatus,
				pxy:              tt.fields.pxy,
				monitor:          tt.fields.monitor,
				handler:          tt.fields.handler,
				msgTransporter:   tt.fields.msgTransporter,
				health:           tt.fields.health,
				lastSendStartMsg: tt.fields.lastSendStartMsg,
				lastStartErr:     tt.fields.lastStartErr,
				closeCh:          tt.fields.closeCh,
				healthNotifyCh:   tt.fields.healthNotifyCh,
				mu:               tt.fields.mu,
				xl:               tt.fields.xl,
				ctx:              tt.fields.ctx,
			}
			if got := pw.GetStatus(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wrapper.GetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
