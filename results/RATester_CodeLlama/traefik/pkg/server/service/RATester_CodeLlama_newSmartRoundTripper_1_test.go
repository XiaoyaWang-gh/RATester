package service

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestNewSmartRoundTripper_1(t *testing.T) {
	type args struct {
		transport          *http.Transport
		forwardingTimeouts *dynamic.ForwardingTimeouts
	}
	tests := []struct {
		name    string
		args    args
		want    *smartRoundTripper
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				transport: &http.Transport{},
			},
			want: &smartRoundTripper{
				http2: &http.Transport{},
				http:  &http.Transport{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := newSmartRoundTripper(tt.args.transport, tt.args.forwardingTimeouts)
			if (err != nil) != tt.wantErr {
				t.Errorf("newSmartRoundTripper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSmartRoundTripper() got = %v, want %v", got, tt.want)
			}
		})
	}
}
