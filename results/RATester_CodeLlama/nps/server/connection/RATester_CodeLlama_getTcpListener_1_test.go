package connection

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestGetTcpListener_1(t *testing.T) {
	type args struct {
		ip string
		p  string
	}
	tests := []struct {
		name    string
		args    args
		want    net.Listener
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				ip: "127.0.0.1",
				p:  "8080",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				ip: "127.0.0.1",
				p:  "8080",
			},
			want:    nil,
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

			got, err := getTcpListener(tt.args.ip, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTcpListener() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTcpListener() = %v, want %v", got, tt.want)
			}
		})
	}
}
