package config

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/msg"
)

func TestNewProxyConfigurerFromMsg_1(t *testing.T) {
	type args struct {
		m         *msg.NewProxy
		serverCfg *v1.ServerConfig
	}
	tests := []struct {
		name    string
		args    args
		want    v1.ProxyConfigurer
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

			got, err := NewProxyConfigurerFromMsg(tt.args.m, tt.args.serverCfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewProxyConfigurerFromMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProxyConfigurerFromMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}
