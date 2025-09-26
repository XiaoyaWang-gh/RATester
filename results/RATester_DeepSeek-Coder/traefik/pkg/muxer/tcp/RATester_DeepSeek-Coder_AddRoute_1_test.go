package tcp

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/tcp"
)

func TestAddRoute_1(t *testing.T) {
	type args struct {
		rule     string
		syntax   string
		priority int
		handler  tcp.Handler
	}
	tests := []struct {
		name    string
		m       *Muxer
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

			m := &Muxer{}
			if err := m.AddRoute(tt.args.rule, tt.args.syntax, tt.args.priority, tt.args.handler); (err != nil) != tt.wantErr {
				t.Errorf("Muxer.AddRoute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
