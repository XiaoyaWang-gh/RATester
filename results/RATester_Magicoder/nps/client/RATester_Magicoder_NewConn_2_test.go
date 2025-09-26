package client

import (
	"fmt"
	"reflect"
	"testing"

	"ehang.io/nps/lib/conn"
)

func TestNewConn_2(t *testing.T) {
	type args struct {
		tp       string
		vkey     string
		server   string
		connType string
		proxyUrl string
	}
	tests := []struct {
		name    string
		args    args
		want    *conn.Conn
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

			got, err := NewConn(tt.args.tp, tt.args.vkey, tt.args.server, tt.args.connType, tt.args.proxyUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConn() = %v, want %v", got, tt.want)
			}
		})
	}
}
