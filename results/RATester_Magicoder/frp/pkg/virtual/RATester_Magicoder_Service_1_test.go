package virtual

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/client"
	netpkg "github.com/fatedier/frp/pkg/util/net"
)

func TestService_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{
		l:   &netpkg.InternalListener{},
		svr: &client.Service{},
	}

	expected := client.svr
	actual := client.Service()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
