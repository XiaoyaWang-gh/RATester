package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewListener_1(t *testing.T) {
	type args struct {
		address string
		network string
	}
	tests := []struct {
		name    string
		args    args
		want    *tcpKeepAliveListener
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				address: "localhost:8080",
				network: "tcp",
			},
			want:    &tcpKeepAliveListener{},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				address: "localhost:8080",
				network: "tcp4",
			},
			want:    &tcpKeepAliveListener{},
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				address: "localhost:8080",
				network: "tcp6",
			},
			want:    &tcpKeepAliveListener{},
			wantErr: false,
		},
		{
			name: "Test case 4",
			args: args{
				address: "localhost:8080",
				network: "invalid",
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

			got, err := newListener(tt.args.address, tt.args.network)
			if (err != nil) != tt.wantErr {
				t.Errorf("newListener() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newListener() = %v, want %v", got, tt.want)
			}
		})
	}
}
