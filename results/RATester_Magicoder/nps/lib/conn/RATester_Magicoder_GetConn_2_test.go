package conn

import (
	"fmt"
	"io"
	"net"
	"reflect"
	"testing"

	"ehang.io/nps/lib/rate"
)

func TestGetConn_2(t *testing.T) {
	type args struct {
		conn     net.Conn
		cpt      bool
		snappy   bool
		rt       *rate.Rate
		isServer bool
	}
	tests := []struct {
		name string
		args args
		want io.ReadWriteCloser
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

			if got := GetConn(tt.args.conn, tt.args.cpt, tt.args.snappy, tt.args.rt, tt.args.isServer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConn() = %v, want %v", got, tt.want)
			}
		})
	}
}
