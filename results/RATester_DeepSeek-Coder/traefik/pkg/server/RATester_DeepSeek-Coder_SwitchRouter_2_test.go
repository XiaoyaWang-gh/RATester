package server

import (
	"fmt"
	"net"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/static"
	tcprouter "github.com/traefik/traefik/v3/pkg/server/router/tcp"
	"github.com/traefik/traefik/v3/pkg/tcp"
)

func TestSwitchRouter_2(t *testing.T) {
	type fields struct {
		listener               net.Listener
		switcher               *tcp.HandlerSwitcher
		transportConfiguration *static.EntryPointsTransport
		tracker                *connectionTracker
		httpServer             *httpServer
		httpsServer            *httpServer
		http3Server            *http3server
	}
	type args struct {
		rt *tcprouter.Router
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

			e := &TCPEntryPoint{
				listener:               tt.fields.listener,
				switcher:               tt.fields.switcher,
				transportConfiguration: tt.fields.transportConfiguration,
				tracker:                tt.fields.tracker,
				httpServer:             tt.fields.httpServer,
				httpsServer:            tt.fields.httpsServer,
				http3Server:            tt.fields.http3Server,
			}
			e.SwitchRouter(tt.args.rt)
		})
	}
}
