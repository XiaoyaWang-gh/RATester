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
		// TODO: Add test cases.
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
				t.Errorf("waitDetectMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
