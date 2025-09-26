package server

import (
	"fmt"
	"net"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/tcp"
)

func TestWriteCloser_1(t *testing.T) {
	type args struct {
		conn net.Conn
	}
	tests := []struct {
		name    string
		args    args
		want    tcp.WriteCloser
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

			got, err := writeCloser(tt.args.conn)
			if (err != nil) != tt.wantErr {
				t.Errorf("writeCloser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("writeCloser() = %v, want %v", got, tt.want)
			}
		})
	}
}
