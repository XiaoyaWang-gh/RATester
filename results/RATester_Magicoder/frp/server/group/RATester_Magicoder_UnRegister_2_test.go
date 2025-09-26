package group

import (
	"fmt"
	"net"
	"testing"

	"github.com/fatedier/frp/pkg/util/vhost"
)

func TestUnRegister_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	g := &HTTPGroup{
		createFuncs: map[string]vhost.CreateConnFunc{
			"proxy1": func(remoteAddr string) (net.Conn, error) { return nil, nil },
			"proxy2": func(remoteAddr string) (net.Conn, error) { return nil, nil },
		},
		pxyNames: []string{"proxy1", "proxy2"},
	}

	g.UnRegister("proxy1")

	if len(g.createFuncs) != 1 {
		t.Errorf("Expected createFuncs length to be 1, but got %d", len(g.createFuncs))
	}

	if len(g.pxyNames) != 1 {
		t.Errorf("Expected pxyNames length to be 1, but got %d", len(g.pxyNames))
	}

	if g.pxyNames[0] != "proxy2" {
		t.Errorf("Expected pxyNames[0] to be 'proxy2', but got %s", g.pxyNames[0])
	}
}
