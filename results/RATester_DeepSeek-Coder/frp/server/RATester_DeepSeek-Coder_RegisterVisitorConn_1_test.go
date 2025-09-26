package server

import (
	"fmt"
	"net"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
	"github.com/fatedier/frp/server/controller"
)

func TestRegisterVisitorConn_1(t *testing.T) {
	ctlManager := &ControlManager{}
	rc := &controller.ResourceController{}
	svr := &Service{
		ctlManager: ctlManager,
		rc:         rc,
	}

	testCases := []struct {
		name        string
		visitorConn net.Conn
		newMsg      *msg.NewVisitorConn
		wantErr     bool
	}{
		{
			name:        "Test case 1",
			visitorConn: nil,
			newMsg: &msg.NewVisitorConn{
				RunID:          "",
				ProxyName:      "",
				SignKey:        "",
				Timestamp:      0,
				UseEncryption:  false,
				UseCompression: false,
			},
			wantErr: true,
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := svr.RegisterVisitorConn(tc.visitorConn, tc.newMsg)
			if (err != nil) != tc.wantErr {
				t.Errorf("RegisterVisitorConn() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
