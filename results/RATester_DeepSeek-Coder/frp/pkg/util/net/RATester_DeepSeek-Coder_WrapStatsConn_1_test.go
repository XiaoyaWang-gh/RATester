package net

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestWrapStatsConn_1(t *testing.T) {
	type args struct {
		conn      net.Conn
		statsFunc func(totalRead, totalWrite int64)
	}
	tests := []struct {
		name string
		args args
		want *StatsConn
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

			if got := WrapStatsConn(tt.args.conn, tt.args.statsFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WrapStatsConn() = %v, want %v", got, tt.want)
			}
		})
	}
}
