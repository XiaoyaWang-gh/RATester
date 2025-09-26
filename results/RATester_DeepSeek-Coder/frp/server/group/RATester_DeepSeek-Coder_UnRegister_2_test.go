package group

import (
	"fmt"
	"net"
	"testing"

	"github.com/fatedier/frp/pkg/util/vhost"
)

func TestUnRegister_2(t *testing.T) {
	g := &HTTPGroup{
		createFuncs: make(map[string]vhost.CreateConnFunc),
		pxyNames:    []string{"proxy1", "proxy2", "proxy3"},
	}

	g.createFuncs["proxy1"] = func(string) (net.Conn, error) { return nil, nil }
	g.createFuncs["proxy2"] = func(string) (net.Conn, error) { return nil, nil }
	g.createFuncs["proxy3"] = func(string) (net.Conn, error) { return nil, nil }

	tests := []struct {
		name        string
		proxyName   string
		wantIsEmpty bool
	}{
		{"unregister existing proxy", "proxy1", false},
		{"unregister non-existing proxy", "proxy4", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			isEmpty := g.UnRegister(tt.proxyName)
			if isEmpty != tt.wantIsEmpty {
				t.Errorf("UnRegister() = %v, want %v", isEmpty, tt.wantIsEmpty)
			}
		})
	}
}
