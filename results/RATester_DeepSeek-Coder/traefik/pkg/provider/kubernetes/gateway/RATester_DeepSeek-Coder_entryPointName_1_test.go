package gateway

import (
	"fmt"
	"testing"

	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

func TestEntryPointName_1(t *testing.T) {
	type args struct {
		port     gatev1.PortNumber
		protocol gatev1.ProtocolType
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				port:     80,
				protocol: "HTTP",
			},
			want:    "entryPoint1",
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				port:     443,
				protocol: "HTTPS",
			},
			want:    "entryPoint2",
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				port:     8080,
				protocol: "TCP",
			},
			want:    "",
			wantErr: true,
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
						Address:        ":80",
						HasHTTPTLSConf: false,
					},
					"entryPoint2": {
						Address:        ":443",
						HasHTTPTLSConf: true,
					},
				},
			}
			got, err := p.entryPointName(tt.args.port, tt.args.protocol)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.entryPointName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Provider.entryPointName() = %v, want %v", got, tt.want)
			}
		})
	}
}
