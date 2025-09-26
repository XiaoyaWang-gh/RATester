package nathole

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDiscoverFromStunServer_1(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				addr: "stun.example.com",
			},
			want:    []string{"192.168.1.1", "192.168.1.2"},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				addr: "stun.example.com",
			},
			want:    []string{"192.168.1.1"},
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				addr: "stun.example.com",
			},
			want:    []string{},
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

			c := &discoverConn{}
			got, err := c.discoverFromStunServer(tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("discoverFromStunServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("discoverFromStunServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
