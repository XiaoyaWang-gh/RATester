package rate

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestNewRateConn_1(t *testing.T) {
	type args struct {
		conn io.ReadWriteCloser
		rate *Rate
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

			if got := NewRateConn(tt.args.conn, tt.args.rate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRateConn() = %v, want %v", got, tt.want)
			}
		})
	}
}
