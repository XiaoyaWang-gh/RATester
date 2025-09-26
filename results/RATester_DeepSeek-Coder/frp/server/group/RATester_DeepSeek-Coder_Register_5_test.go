package group

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/util/vhost"
)

func TestRegister_5(t *testing.T) {
	type args struct {
		proxyName   string
		group       string
		groupKey    string
		routeConfig vhost.RouteConfig
	}
	tests := []struct {
		name    string
		g       *HTTPGroup
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

			g := &HTTPGroup{
				group:           tt.g.group,
				groupKey:        tt.g.groupKey,
				domain:          tt.g.domain,
				location:        tt.g.location,
				routeByHTTPUser: tt.g.routeByHTTPUser,
				createFuncs:     tt.g.createFuncs,
				pxyNames:        tt.g.pxyNames,
				ctl:             tt.g.ctl,
			}
			if err := g.Register(tt.args.proxyName, tt.args.group, tt.args.groupKey, tt.args.routeConfig); (err != nil) != tt.wantErr {
				t.Errorf("HTTPGroup.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
