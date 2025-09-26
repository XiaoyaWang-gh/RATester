package consulcatalog

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestServersTransport_1(t *testing.T) {
	type args struct {
		item itemData
	}
	tests := []struct {
		name string
		c    *connectCert
		args args
		want *dynamic.ServersTransport
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

			if got := tt.c.serversTransport(tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connectCert.serversTransport() = %v, want %v", got, tt.want)
			}
		})
	}
}
