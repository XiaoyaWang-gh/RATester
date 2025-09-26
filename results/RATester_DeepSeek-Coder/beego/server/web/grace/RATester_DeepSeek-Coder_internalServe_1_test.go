package grace

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"testing"
)

func TestInternalServe_1(t *testing.T) {
	type fields struct {
		Server            *http.Server
		ln                net.Listener
		SignalHooks       map[int]map[os.Signal][]func()
		sigChan           chan os.Signal
		isChild           bool
		state             uint8
		Network           string
		terminalChan      chan error
		shutdownCallbacks []func()
	}
	type args struct {
		ln net.Listener
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
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

			srv := &Server{
				Server:            tt.fields.Server,
				ln:                tt.fields.ln,
				SignalHooks:       tt.fields.SignalHooks,
				sigChan:           tt.fields.sigChan,
				isChild:           tt.fields.isChild,
				state:             tt.fields.state,
				Network:           tt.fields.Network,
				terminalChan:      tt.fields.terminalChan,
				shutdownCallbacks: tt.fields.shutdownCallbacks,
			}
			if err := srv.internalServe(tt.args.ln); (err != nil) != tt.wantErr {
				t.Errorf("Server.internalServe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
