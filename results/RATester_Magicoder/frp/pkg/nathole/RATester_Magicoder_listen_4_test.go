package nathole

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestListen_4(t *testing.T) {
	type args struct {
		localAddr string
	}
	tests := []struct {
		name    string
		args    args
		want    *discoverConn
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				localAddr: "127.0.0.1:8080",
			},
			want: &discoverConn{
				conn:        &net.UDPConn{},
				localAddr:   &net.UDPAddr{},
				messageChan: make(chan *Message, 10),
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				localAddr: "",
			},
			want: &discoverConn{
				conn:        &net.UDPConn{},
				localAddr:   &net.UDPAddr{},
				messageChan: make(chan *Message, 10),
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				localAddr: "invalid_address",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := listen(tt.args.localAddr)
			if (err != nil) != tt.wantErr {
				t.Errorf("listen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listen() = %v, want %v", got, tt.want)
			}
		})
	}
}
