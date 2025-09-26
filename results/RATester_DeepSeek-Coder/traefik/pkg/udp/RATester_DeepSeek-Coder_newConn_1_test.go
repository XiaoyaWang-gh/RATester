package udp

import (
	"fmt"
	"net"
	"reflect"
	"testing"
	"time"
)

func TestNewConn_1(t *testing.T) {
	type args struct {
		rAddr net.Addr
	}
	tests := []struct {
		name string
		args args
		want *Conn
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

			l := &Listener{
				accepting: true,
				timeout:   30 * time.Second,
			}
			if got := l.newConn(tt.args.rAddr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Listener.newConn() = %v, want %v", got, tt.want)
			}
		})
	}
}
