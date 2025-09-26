package proxy

import (
	"fmt"
	"net/http"
	"testing"

	"ehang.io/nps/lib/conn"
)

func TesthandleHttp_1(t *testing.T) {
	type args struct {
		c *conn.Conn
		r *http.Request
	}
	tests := []struct {
		name string
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

			s := &httpServer{}
			s.handleHttp(tt.args.c, tt.args.r)
		})
	}
}
