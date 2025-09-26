package group

import (
	"fmt"
	"sync"
	"testing"

	"github.com/fatedier/frp/pkg/util/vhost"
)

func TestChooseEndpoint_1(t *testing.T) {
	type fields struct {
		group           string
		groupKey        string
		domain          string
		location        string
		routeByHTTPUser string
		createFuncs     map[string]vhost.CreateConnFunc
		pxyNames        []string
		index           uint64
		ctl             *HTTPGroupController
		mu              sync.RWMutex
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
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
				group:           tt.fields.group,
				groupKey:        tt.fields.groupKey,
				domain:          tt.fields.domain,
				location:        tt.fields.location,
				routeByHTTPUser: tt.fields.routeByHTTPUser,
				createFuncs:     tt.fields.createFuncs,
				pxyNames:        tt.fields.pxyNames,
				index:           tt.fields.index,
				ctl:             tt.fields.ctl,
				mu:              tt.fields.mu,
			}
			got, err := g.chooseEndpoint()
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPGroup.chooseEndpoint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HTTPGroup.chooseEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
