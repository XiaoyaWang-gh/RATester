package tcp

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestBuildSpiffeAuthorizer_1(t *testing.T) {
	type args struct {
		cfg *dynamic.Spiffe
	}
	tests := []struct {
		name    string
		args    args
		want    tlsconfig.Authorizer
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

			got, err := buildSpiffeAuthorizer(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildSpiffeAuthorizer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildSpiffeAuthorizer() = %v, want %v", got, tt.want)
			}
		})
	}
}
