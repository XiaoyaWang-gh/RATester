package gateway

import (
	"fmt"
	"testing"

	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

func TestEntryPointName_1(t *testing.T) {
	tests := []struct {
		name     string
		port     gatev1.PortNumber
		protocol gatev1.ProtocolType
		want     string
		wantErr  bool
	}{
		{
			name:     "test case 1",
			port:     80,
			protocol: gatev1.HTTPProtocolType,
			want:     "entryPoint1",
			wantErr:  false,
		},
		{
			name:     "test case 2",
			port:     443,
			protocol: gatev1.HTTPSProtocolType,
			want:     "entryPoint2",
			wantErr:  false,
		},
		{
			name:     "test case 3",
			port:     8080,
			protocol: gatev1.HTTPProtocolType,
			want:     "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &Provider{
				EntryPoints: map[string]Entrypoint{
					"entryPoint1": {
						Address:        "127.0.0.1:80",
						HasHTTPTLSConf: false,
					},
					"entryPoint2": {
						Address:        "127.0.0.1:443",
						HasHTTPTLSConf: true,
					},
				},
			}
			got, err := p.entryPointName(tt.port, tt.protocol)
			if (err != nil) != tt.wantErr {
				t.Errorf("entryPointName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("entryPointName() = %v, want %v", got, tt.want)
			}
		})
	}
}
