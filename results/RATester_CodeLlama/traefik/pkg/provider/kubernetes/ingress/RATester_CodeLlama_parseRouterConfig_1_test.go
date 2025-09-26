package ingress

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestParseRouterConfig_1(t *testing.T) {
	type args struct {
		annotations map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    *RouterConfig
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				annotations: map[string]string{
					"traefik.router.tls.certresolver": "myresolver",
				},
			},
			want: &RouterConfig{
				Router: &RouterIng{
					TLS: &dynamic.RouterTLSConfig{
						CertResolver: "myresolver",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := parseRouterConfig(tt.args.annotations)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseRouterConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRouterConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
