package file

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/safe"
)

func TestAddWatcher_1(t *testing.T) {
	type args struct {
		pool              *safe.Pool
		items             []string
		configurationChan chan<- dynamic.Message
		callback          func(chan<- dynamic.Message) error
	}
	tests := []struct {
		name    string
		args    args
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

			p := &Provider{}
			if err := p.addWatcher(tt.args.pool, tt.args.items, tt.args.configurationChan, tt.args.callback); (err != nil) != tt.wantErr {
				t.Errorf("Provider.addWatcher() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
