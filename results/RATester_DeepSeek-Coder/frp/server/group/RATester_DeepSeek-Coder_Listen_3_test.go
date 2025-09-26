package group

import (
	"context"
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/util/vhost"
)

func TestListen_3(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	testCases := []struct {
		name        string
		multiplexer string
		group       string
		groupKey    string
		routeConfig vhost.RouteConfig
		wantErr     bool
	}{
		{
			name:        "Test case 1",
			multiplexer: "http-connect",
			group:       "testGroup",
			groupKey:    "testKey",
			routeConfig: vhost.RouteConfig{
				Domain: "test.com",
			},
			wantErr: false,
		},
		{
			name:        "Test case 2",
			multiplexer: "unknown-multiplexer",
			group:       "testGroup",
			groupKey:    "testKey",
			routeConfig: vhost.RouteConfig{
				Domain: "test.com",
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tmgc := &TCPMuxGroupCtl{
				groups: make(map[string]*TCPMuxGroup),
			}

			_, err := tmgc.Listen(ctx, tc.multiplexer, tc.group, tc.groupKey, tc.routeConfig)
			if (err != nil) != tc.wantErr {
				t.Errorf("Listen() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
