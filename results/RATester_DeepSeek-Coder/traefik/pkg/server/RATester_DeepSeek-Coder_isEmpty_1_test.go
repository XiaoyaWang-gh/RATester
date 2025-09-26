package server

import (
	"fmt"
	"net"
	"sync"
	"testing"
)

func TestIsEmpty_1(t *testing.T) {
	type fields struct {
		connsMu sync.RWMutex
		conns   map[net.Conn]struct{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Test isEmpty when map is empty",
			fields: fields{
				connsMu: sync.RWMutex{},
				conns:   make(map[net.Conn]struct{}),
			},
			want: true,
		},
		{
			name: "Test isEmpty when map is not empty",
			fields: fields{
				connsMu: sync.RWMutex{},
				conns: map[net.Conn]struct{}{
					&net.TCPConn{}: {},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &connectionTracker{
				connsMu: tt.fields.connsMu,
				conns:   tt.fields.conns,
			}
			if got := c.isEmpty(); got != tt.want {
				t.Errorf("connectionTracker.isEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
