package proxy

import (
	"fmt"
	"testing"

	"ehang.io/nps/lib/conn"
	"ehang.io/nps/lib/file"
)

func TestDealClient_1(t *testing.T) {
	type args struct {
		c          *conn.Conn
		client     *file.Client
		addr       string
		rb         []byte
		tp         string
		f          func()
		flow       *file.Flow
		localProxy bool
	}
	tests := []struct {
		name    string
		s       *BaseServer
		args    args
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

			if err := tt.s.DealClient(tt.args.c, tt.args.client, tt.args.addr, tt.args.rb, tt.args.tp, tt.args.f, tt.args.flow, tt.args.localProxy); (err != nil) != tt.wantErr {
				t.Errorf("BaseServer.DealClient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
