package v1

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
)

func TestMarshalToMsg_1(t *testing.T) {
	type fields struct {
		ProxyBaseConfig ProxyBaseConfig
		Secretkey       string
		AllowUsers      []string
	}
	type args struct {
		m *msg.NewProxy
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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

			c := &SUDPProxyConfig{
				ProxyBaseConfig: tt.fields.ProxyBaseConfig,
				Secretkey:       tt.fields.Secretkey,
				AllowUsers:      tt.fields.AllowUsers,
			}
			c.MarshalToMsg(tt.args.m)
		})
	}
}
