package proxy

import (
	"fmt"
	"testing"
)

func TestStartProxy_1(t *testing.T) {
	pm := &Manager{
		proxies: map[string]*Wrapper{
			"testProxy": &Wrapper{
				WorkingStatus: WorkingStatus{
					Name: "testProxy",
				},
			},
		},
	}

	tests := []struct {
		name          string
		proxyName     string
		remoteAddr    string
		serverRespErr string
		wantErr       bool
	}{
		{
			name:          "Test case 1: Proxy found",
			proxyName:     "testProxy",
			remoteAddr:    "127.0.0.1:8080",
			serverRespErr: "",
			wantErr:       false,
		},
		{
			name:          "Test case 2: Proxy not found",
			proxyName:     "nonExistentProxy",
			remoteAddr:    "127.0.0.1:8080",
			serverRespErr: "",
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := pm.StartProxy(tt.proxyName, tt.remoteAddr, tt.serverRespErr)
			if (err != nil) != tt.wantErr {
				t.Errorf("StartProxy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
