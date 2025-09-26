package vhost

import (
	"fmt"
	"net"
	"testing"
)

func TestHandle_2(t *testing.T) {
	type args struct {
		c net.Conn
	}
	tests := []struct {
		name string
		v    *Muxer
		args args
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

			tt.v.handle(tt.args.c)
		})
	}
}
