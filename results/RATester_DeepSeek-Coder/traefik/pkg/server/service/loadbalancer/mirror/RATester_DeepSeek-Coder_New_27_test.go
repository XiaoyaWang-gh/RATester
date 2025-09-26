package mirror

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/safe"
)

func TestNew_27(t *testing.T) {
	type args struct {
		handler     http.Handler
		pool        *safe.Pool
		mirrorBody  bool
		maxBodySize int64
		hc          *dynamic.HealthCheck
	}
	tests := []struct {
		name string
		args args
		want *Mirroring
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

			got := New(tt.args.handler, tt.args.pool, tt.args.mirrorBody, tt.args.maxBodySize, tt.args.hc)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
