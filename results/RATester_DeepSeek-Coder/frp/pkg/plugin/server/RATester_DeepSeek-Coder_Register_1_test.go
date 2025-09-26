package plugin

import (
	"fmt"
	"testing"
)

func TestRegister_1(t *testing.T) {
	type fields struct {
		loginPlugins       []Plugin
		newProxyPlugins    []Plugin
		closeProxyPlugins  []Plugin
		pingPlugins        []Plugin
		newWorkConnPlugins []Plugin
		newUserConnPlugins []Plugin
	}
	type args struct {
		p Plugin
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

			m := &Manager{
				loginPlugins:       tt.fields.loginPlugins,
				newProxyPlugins:    tt.fields.newProxyPlugins,
				closeProxyPlugins:  tt.fields.closeProxyPlugins,
				pingPlugins:        tt.fields.pingPlugins,
				newWorkConnPlugins: tt.fields.newWorkConnPlugins,
				newUserConnPlugins: tt.fields.newUserConnPlugins,
			}
			m.Register(tt.args.p)
		})
	}
}
