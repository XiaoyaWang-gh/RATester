package server

import (
	"fmt"
	"testing"

	"github.com/rs/zerolog"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestLogConfiguration_1(t *testing.T) {
	type args struct {
		logger    zerolog.Logger
		configMsg dynamic.Message
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

			logConfiguration(tt.args.logger, tt.args.configMsg)
		})
	}
}
