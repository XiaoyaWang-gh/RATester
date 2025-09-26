package pmux

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestGetHttpListener_2(t *testing.T) {
	type fields struct {
		Listener    net.Listener
		port        int
		isClose     bool
		managerHost string
		clientConn  chan *PortConn
		httpConn    chan *PortConn
		httpsConn   chan *PortConn
		managerConn chan *PortConn
	}
	tests := []struct {
		name   string
		fields fields
		want   net.Listener
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

			pMux := &PortMux{
				Listener:    tt.fields.Listener,
				port:        tt.fields.port,
				isClose:     tt.fields.isClose,
				managerHost: tt.fields.managerHost,
				clientConn:  tt.fields.clientConn,
				httpConn:    tt.fields.httpConn,
				httpsConn:   tt.fields.httpsConn,
				managerConn: tt.fields.managerConn,
			}
			if got := pMux.GetHttpListener(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PortMux.GetHttpListener() = %v, want %v", got, tt.want)
			}
		})
	}
}
