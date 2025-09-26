package httplib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewClient_1(t *testing.T) {
	type args struct {
		name     string
		endpoint string
		opts     []ClientOption
	}
	tests := []struct {
		name    string
		args    args
		want    *Client
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				name:     "test",
				endpoint: "http://127.0.0.1:8080",
				opts: []ClientOption{
					func(client *Client) {
						client.CommonOpts = append(client.CommonOpts, WithHeader("Content-Type", "application/json"))
					},
				},
			},
			want: &Client{
				Name:       "test",
				Endpoint:   "http://127.0.0.1:8080",
				CommonOpts: []BeegoHTTPRequestOption{WithHeader("Content-Type", "application/json")},
				Setting:    GetDefaultSetting(),
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

			got, err := NewClient(tt.args.name, tt.args.endpoint, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}
