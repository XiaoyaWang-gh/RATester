package conn

import (
	"fmt"
	"net"
	"testing"

	"ehang.io/nps/lib/file"
	"ehang.io/nps/lib/rate"
)

func TestCopyWaitGroup_1(t *testing.T) {
	type args struct {
		conn1    net.Conn
		conn2    net.Conn
		crypt    bool
		snappy   bool
		rate     *rate.Rate
		flow     *file.Flow
		isServer bool
		rb       []byte
	}
	tests := []struct {
		name string
		args args
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

			CopyWaitGroup(tt.args.conn1, tt.args.conn2, tt.args.crypt, tt.args.snappy, tt.args.rate, tt.args.flow, tt.args.isServer, tt.args.rb)
		})
	}
}
