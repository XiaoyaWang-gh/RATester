package virtual

import (
	"fmt"
	"testing"

	netpkg "github.com/fatedier/frp/pkg/util/net"
)

func TestConnect_2(t *testing.T) {
	t.Parallel()

	pc := &pipeConnector{
		peerListener: &netpkg.InternalListener{},
	}

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Test case 1",
			wantErr: false,
		},
		{
			name:    "Test case 2",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := pc.Connect()
			if (err != nil) != tt.wantErr {
				t.Errorf("pipeConnector.Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
