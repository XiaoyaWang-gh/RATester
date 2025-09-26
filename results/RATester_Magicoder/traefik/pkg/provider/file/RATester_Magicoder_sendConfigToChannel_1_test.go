package file

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestsendConfigToChannel_1(t *testing.T) {
	type args struct {
		configurationChan chan<- dynamic.Message
		configuration     *dynamic.Configuration
	}
	tests := []struct {
		name string
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

			sendConfigToChannel(tt.args.configurationChan, tt.args.configuration)
		})
	}
}
