package goroutine

import (
	"fmt"
	"io"
	"net"
	"reflect"
	"sync"
	"testing"

	"ehang.io/nps/lib/file"
)

func TestNewConns_1(t *testing.T) {
	type args struct {
		c1   io.ReadWriteCloser
		c2   net.Conn
		flow *file.Flow
		wg   *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
		want Conns
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

			if got := NewConns(tt.args.c1, tt.args.c2, tt.args.flow, tt.args.wg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConns() = %v, want %v", got, tt.want)
			}
		})
	}
}
