package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew_5(t *testing.T) {
	type args struct {
		host string
		port int
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "test1",
			args: args{
				host: "127.0.0.1",
				port: 7000,
			},
			want: &Client{
				address: "127.0.0.1:7000",
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

			if got := New(tt.args.host, tt.args.port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
