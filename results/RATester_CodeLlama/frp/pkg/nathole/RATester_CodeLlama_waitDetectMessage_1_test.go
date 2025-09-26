package nathole

import (
	"context"
	"fmt"
	"net"
	"reflect"
	"testing"
	"time"
)

func TestWaitDetectMessage_1(t *testing.T) {
	type args struct {
		ctx     context.Context
		conn    *net.UDPConn
		sid     string
		key     []byte
		timeout time.Duration
		role    string
	}
	tests := []struct {
		name    string
		args    args
		want    *net.UDPAddr
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				ctx:     context.Background(),
				conn:    &net.UDPConn{},
				sid:     "sid",
				key:     []byte("key"),
				timeout: time.Second,
				role:    "role",
			},
			want:    &net.UDPAddr{},
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

			got, err := waitDetectMessage(tt.args.ctx, tt.args.conn, tt.args.sid, tt.args.key, tt.args.timeout, tt.args.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("waitDetectMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("waitDetectMessage() got = %v, want %v", got, tt.want)
			}
		})
	}
}
