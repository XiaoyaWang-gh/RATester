package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/static"
	"github.com/traefik/traefik/v3/pkg/udp"
)

func TestStart_3(t *testing.T) {
	type fields struct {
		listener               *udp.Listener
		switcher               *udp.HandlerSwitcher
		transportConfiguration *static.EntryPointsTransport
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

			ep := &UDPEntryPoint{
				listener:               tt.fields.listener,
				switcher:               tt.fields.switcher,
				transportConfiguration: tt.fields.transportConfiguration,
			}
			ep.Start(tt.args.ctx)
		})
	}
}
