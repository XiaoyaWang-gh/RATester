package acme

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestSetConfigListenerChan_1(t *testing.T) {
	type args struct {
		configFromListenerChan chan dynamic.Configuration
	}
	tests := []struct {
		name string
		p    *Provider
		args args
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

			tt.p.SetConfigListenerChan(tt.args.configFromListenerChan)
		})
	}
}
