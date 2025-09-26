package acme

import (
	"fmt"
	"testing"

	traefiktls "github.com/traefik/traefik/v3/pkg/tls"
)

func TestSetTLSManager_1(t *testing.T) {
	testCases := []struct {
		name    string
		manager *traefiktls.Manager
		wantErr bool
	}{
		{
			name:    "set tls manager",
			manager: &traefiktls.Manager{},
			wantErr: false,
		},
		{
			name:    "set nil tls manager",
			manager: nil,
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &Provider{}
			p.SetTLSManager(tc.manager)
			if (p.tlsManager == nil) != tc.wantErr {
				t.Errorf("SetTLSManager() error = %v, wantErr %v", p.tlsManager == nil, tc.wantErr)
			}
		})
	}
}
