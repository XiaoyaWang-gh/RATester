package tcp

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/tcp"
)

func TestGetConn_1(t *testing.T) {
	type args struct {
		conn   tcp.WriteCloser
		peeked string
	}
	tests := []struct {
		name string
		args args
		want tcp.WriteCloser
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

			r := &Router{}
			if got := r.GetConn(tt.args.conn, tt.args.peeked); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConn() = %v, want %v", got, tt.want)
			}
		})
	}
}
