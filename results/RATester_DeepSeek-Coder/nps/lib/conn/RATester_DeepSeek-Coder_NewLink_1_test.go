package conn

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewLink_1(t *testing.T) {
	type args struct {
		connType   string
		host       string
		crypt      bool
		compress   bool
		remoteAddr string
		localProxy bool
		opts       []Option
	}
	tests := []struct {
		name string
		args args
		want *Link
	}{
		{
			name: "TestNewLink",
			args: args{
				connType:   "tcp",
				host:       "localhost:8080",
				crypt:      true,
				compress:   true,
				remoteAddr: "127.0.0.1:8080",
				localProxy: true,
				opts:       nil,
			},
			want: &Link{
				ConnType:   "tcp",
				Host:       "localhost:8080",
				Crypt:      true,
				Compress:   true,
				LocalProxy: true,
				RemoteAddr: "127.0.0.1:8080",
				Option:     Options{},
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

			if got := NewLink(tt.args.connType, tt.args.host, tt.args.crypt, tt.args.compress, tt.args.remoteAddr, tt.args.localProxy, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLink() = %v, want %v", got, tt.want)
			}
		})
	}
}
