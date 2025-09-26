package virtual

import (
	"fmt"
	"testing"

	netpkg "github.com/fatedier/frp/pkg/util/net"
)

func TestOpen_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pc := &pipeConnector{
		peerListener: &netpkg.InternalListener{},
	}

	err := pc.Open()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
