package service

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/runtime"
	"github.com/traefik/traefik/v3/pkg/safe"
	"github.com/traefik/traefik/v3/pkg/server/middleware"
)

func TestNewManager_7(t *testing.T) {
	type args struct {
		configs             map[string]*runtime.ServiceInfo
		observabilityMgr    *middleware.ObservabilityMgr
		routinePool         *safe.Pool
		roundTripperManager RoundTripperGetter
	}
	tests := []struct {
		name string
		args args
		want *Manager
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

			if got := NewManager(tt.args.configs, tt.args.observabilityMgr, tt.args.routinePool, tt.args.roundTripperManager); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewManager() = %v, want %v", got, tt.want)
			}
		})
	}
}
