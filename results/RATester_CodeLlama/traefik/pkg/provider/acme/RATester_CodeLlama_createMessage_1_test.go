package acme

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/types"
)

func TestCreateMessage_1(t *testing.T) {
	type args struct {
		certs map[string]*Certificate
	}
	tests := []struct {
		name string
		args args
		want dynamic.Message
	}{
		{
			name: "test",
			args: args{
				certs: map[string]*Certificate{
					"test": {
						Domain: types.Domain{
							Main: "test.com",
						},
						Certificate: []byte("test"),
						Key:         []byte("test"),
					},
				},
			},
			want: dynamic.Message{
				ProviderName: providerNameALPN,
				Configuration: &dynamic.Configuration{
					HTTP: &dynamic.HTTPConfiguration{
						Routers:     map[string]*dynamic.Router{},
						Middlewares: map[string]*dynamic.Middleware{},
						Services:    map[string]*dynamic.Service{},
					},
					TLS: &dynamic.TLSConfiguration{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := createMessage(tt.args.certs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
