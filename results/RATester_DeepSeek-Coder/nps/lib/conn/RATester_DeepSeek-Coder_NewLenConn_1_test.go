package conn

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestNewLenConn_1(t *testing.T) {
	type args struct {
		conn io.Writer
	}
	tests := []struct {
		name string
		args args
		want *LenConn
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

			if got := NewLenConn(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLenConn() = %v, want %v", got, tt.want)
			}
		})
	}
}
