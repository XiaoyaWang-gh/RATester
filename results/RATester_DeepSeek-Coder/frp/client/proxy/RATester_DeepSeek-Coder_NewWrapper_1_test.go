package proxy

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/fatedier/frp/client/event"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/transport"
)

func TestNewWrapper_1(t *testing.T) {
	type args struct {
		ctx            context.Context
		cfg            v1.ProxyConfigurer
		clientCfg      *v1.ClientCommonConfig
		eventHandler   event.Handler
		msgTransporter transport.MessageTransporter
	}
	tests := []struct {
		name string
		args args
		want *Wrapper
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

			if got := NewWrapper(tt.args.ctx, tt.args.cfg, tt.args.clientCfg, tt.args.eventHandler, tt.args.msgTransporter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWrapper() = %v, want %v", got, tt.want)
			}
		})
	}
}
