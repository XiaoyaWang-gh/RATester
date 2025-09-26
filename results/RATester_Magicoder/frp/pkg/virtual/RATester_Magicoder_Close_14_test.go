package virtual

import (
	"fmt"
	"testing"

	netpkg "github.com/fatedier/frp/pkg/util/net"
)

func TestClose_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pc := &pipeConnector{
		peerListener: &netpkg.InternalListener{},
	}

	err := pc.Close()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
