package vhost

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestRun_3(t *testing.T) {
	type fields struct {
		listener       net.Listener
		timeout        time.Duration
		vhostFunc      muxFunc
		checkAuth      authFunc
		successHook    successHookFunc
		failHook       failHookFunc
		rewriteHost    hostRewriteFunc
		registryRouter *Routers
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

			v := &Muxer{
				listener:       tt.fields.listener,
				timeout:        tt.fields.timeout,
				vhostFunc:      tt.fields.vhostFunc,
				checkAuth:      tt.fields.checkAuth,
				successHook:    tt.fields.successHook,
				failHook:       tt.fields.failHook,
				rewriteHost:    tt.fields.rewriteHost,
				registryRouter: tt.fields.registryRouter,
			}
			v.run()
		})
	}
}
