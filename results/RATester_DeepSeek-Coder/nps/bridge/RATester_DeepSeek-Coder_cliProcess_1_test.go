package bridge

import (
	"fmt"
	"testing"

	"ehang.io/nps/lib/conn"
)

func TestCliProcess_1(t *testing.T) {
	type args struct {
		c *conn.Conn
	}
	tests := []struct {
		name string
		s    *Bridge
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

			tt.s.cliProcess(tt.args.c)
		})
	}
}
