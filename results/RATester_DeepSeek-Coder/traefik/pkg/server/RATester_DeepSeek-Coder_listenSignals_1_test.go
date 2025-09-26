package server

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/traefik/traefik/v3/pkg/safe"
	"github.com/traefik/traefik/v3/pkg/server/middleware"
)

func TestListenSignals_1(t *testing.T) {
	type fields struct {
		watcher          *ConfigurationWatcher
		tcpEntryPoints   TCPEntryPoints
		udpEntryPoints   UDPEntryPoints
		observabilityMgr *middleware.ObservabilityMgr
		signals          chan os.Signal
		stopChan         chan bool
		routinesPool     *safe.Pool
	}
	type args struct {
		ctx context.Context
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

			s := &Server{
				watcher:          tt.fields.watcher,
				tcpEntryPoints:   tt.fields.tcpEntryPoints,
				udpEntryPoints:   tt.fields.udpEntryPoints,
				observabilityMgr: tt.fields.observabilityMgr,
				signals:          tt.fields.signals,
				stopChan:         tt.fields.stopChan,
				routinesPool:     tt.fields.routinesPool,
			}
			s.listenSignals(tt.args.ctx)
		})
	}
}
