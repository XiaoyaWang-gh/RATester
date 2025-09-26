package udp

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewProxy_1(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    *Proxy
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				address: "localhost:8080",
			},
			want: &Proxy{
				target: "localhost:8080",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				address: "invalid_address",
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

			got, err := NewProxy(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewProxy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}
