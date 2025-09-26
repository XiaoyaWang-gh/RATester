package nathole

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDoSTUNRequest_1(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		args    args
		want    *stunResponse
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				addr: "stun.example.com:3478",
			},
			want: &stunResponse{
				externalAddr: "192.0.2.1:1234",
				otherAddr:    "192.0.2.2:1234",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				addr: "invalid_stun_server",
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

			c := &discoverConn{}
			got, err := c.doSTUNRequest(tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("doSTUNRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doSTUNRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
