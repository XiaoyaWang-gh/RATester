package headers

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestNewSecure_1(t *testing.T) {
	type args struct {
		next       http.Handler
		cfg        dynamic.Headers
		contextKey string
	}
	tests := []struct {
		name string
		args args
		want *secureHeader
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

			if got := newSecure(tt.args.next, tt.args.cfg, tt.args.contextKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSecure() = %v, want %v", got, tt.want)
			}
		})
	}
}
