package ports

import (
	"fmt"
	"sync"
	"testing"
)

func TestIsPortAvailable_1(t *testing.T) {
	type fields struct {
		reservedPorts map[string]*PortCtx
		usedPorts     map[int]*PortCtx
		freePorts     map[int]struct{}
		bindAddr      string
		netType       string
		mu            sync.Mutex
	}
	type args struct {
		port int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
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

			pm := &Manager{
				reservedPorts: tt.fields.reservedPorts,
				usedPorts:     tt.fields.usedPorts,
				freePorts:     tt.fields.freePorts,
				bindAddr:      tt.fields.bindAddr,
				netType:       tt.fields.netType,
				mu:            tt.fields.mu,
			}
			if got := pm.isPortAvailable(tt.args.port); got != tt.want {
				t.Errorf("Manager.isPortAvailable() = %v, want %v", got, tt.want)
			}
		})
	}
}
