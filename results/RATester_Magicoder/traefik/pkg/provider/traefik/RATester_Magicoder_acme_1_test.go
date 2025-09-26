package traefik

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestAcme_1(t *testing.T) {
	type args struct {
		cfg *dynamic.Configuration
	}
	tests := []struct {
		name string
		args args
		want *dynamic.Configuration
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

			i := &Provider{}
			i.acme(tt.args.cfg)
			if !reflect.DeepEqual(tt.args.cfg, tt.want) {
				t.Errorf("acme() = %v, want %v", tt.args.cfg, tt.want)
			}
		})
	}
}
