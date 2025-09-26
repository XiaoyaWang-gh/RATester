package conn

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestNewSnappyConn_1(t *testing.T) {
	type args struct {
		conn io.ReadWriteCloser
	}
	tests := []struct {
		name string
		args args
		want *SnappyConn
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

			if got := NewSnappyConn(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSnappyConn() = %v, want %v", got, tt.want)
			}
		})
	}
}
